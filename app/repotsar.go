package haus

import(
	"io/ioutil"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/SearchSpring/RepoTsar/gitutils"
)

// Signature represents the values to be use in a git signature.
type Signature struct {
	Name string
	Email string
}

// Repo represents a single git repo definition.
type Repo struct {
	URL string
	Path string
	Branch string

}

// CloneRepo clones a git repo, hand it a repo name, returns error.
func (r *Repo) CloneRepo(n string) error {
	fmt.Printf("Cloning repo: %s into %s\n",n,r.Path)
	cloneinfo := &gitutils.CloneInfo{
		Reponame: n,
		Path: r.Path,
		URL: r.URL,
		Branch: r.Branch,
	}
	_,err := cloneinfo.CloneRepo()
	if err != nil {
		return err
	}
	return nil
}

// RepoYml represents a collection of repos configs to be turned into YAML.
type RepoYml struct {
	Signature Signature
	Repos map[string]Repo
}

// AddCfg adds a map[string]Repo to RepoYml.
func (y *RepoYml) AddCfg(r map[string]Repo) {
	for k,v := range r {
		if y.Repos == nil {
			y.Repos = make(map[string]Repo)
		}
		y.Repos[k] = v
	}
}

// WriteYml writes out Cfgs to yaml file.  Hand it full path
// to yaml file you wish to write out. Returns string of yaml text.
func (y *RepoYml) WriteYml(filename string) (string,error) {
	yaml,err := yaml.Marshal(y)
	if err != nil {
		return "",err
	}
	err = ioutil.WriteFile(filename, yaml, 0644)
	if err != nil {
		return "",err
	}
	return string(yaml[:]),err	
}


