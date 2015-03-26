package fileutils

import(
	"os"
	"strings"
	"path/filepath"
	"fmt"
)


// expandTilde expands ~ to value of ENV HOME
func expandTilde(f string) string {
	if strings.HasPrefix(f, "~"+string(filepath.Separator)) {
		return os.Getenv("HOME") + f[1:]
	}
	return f
}

// CreatePath check for path and create if it doesn't exist
func CreatePath(path string) (string, error) {
	// Get a abs path
	absPath,err := filepath.Abs(expandTilde(path))
	if err != nil {
		return "", err
	}
	// check if path exists
	stat,err := os.Stat(absPath)
	// create path if it doesn't or throw error if it's not a dir 
	if stat == nil && os.IsNotExist(err) {
		err := os.MkdirAll(absPath, 0755)
		if err != nil {
			return "", err
		}
	} else if !stat.IsDir() {
		err := fmt.Errorf("path: %#v exists but is not a directory\n", path)
		return "", err
	}

	return absPath, nil
}

// CreateSymlink check for symlink and create it if doesn't exist
func CreateSymlink(opath string, npath string) (string, error) {
	// Get clean abs paths
	cleanOpath := filepath.Clean(expandTilde(opath))
	cleanNpath := filepath.Clean(expandTilde(npath))
	stat,err := os.Lstat(cleanNpath)
	if stat == nil {
		// Create dir path for npath if it doesn't exist
		dirNpath := filepath.Dir(cleanNpath)
		err = nil
		_,err = CreatePath(dirNpath)
		if err != nil {
			return "", fmt.Errorf("Cannot create dir path \n%#v\n for \n%#v\n : %s", dirNpath, cleanNpath, err)
		}
		// Symlink
		err = os.Symlink(cleanOpath+"/",cleanNpath)
		if err != nil {
			return "",err
		}
	}
	return cleanNpath, err
}

