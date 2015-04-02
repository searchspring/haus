package haus

import(
	"os"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"github.com/SearchSpring/RepoTsar/gitutils"
)

// Config represents the configurations in haus config file.
type Config struct {
	Name string
	Email string
	Path string
	Pwd string
	Hausrepo string
	Environments map[string]Environment
}

// ReadConfig reads the config file from the supplied full path and
// returns a Config and error.
func ReadConfig(filename string, usrcfgfile string, branch string )(*Config, error) {
	config := &Config{}

	// If the configfile is missing, try to check it out from git repo
	_,err := os.Stat(filename)
	if err != nil {
		// Get the url for the git repo from user config
		err = readCfg(usrcfgfile,config)
		if err != nil {
			return config,err
		}
		// If the url is defined, clone the repo
		if config.Hausrepo != "" {
			cloneinfo := &gitutils.CloneInfo{
				Reponame: "hauscfg",
				Path: ".",
				URL: config.Hausrepo,
				Branch: branch,
			}
			fmt.Printf("Cloning repo hauscfg from %s\n", config.Hausrepo)
			_,err = cloneinfo.CloneRepo()
			if err != nil {
				return config, err
			}
		} else {
			// There's no haus yaml, and hausrepo isn't defined in the user config
			err = fmt.Errorf("No %s file and %s missing 'hausrepo'.", filename,usrcfgfile)
			return config, err
		}
	} 

	// Read haus yaml file
	err = readCfg(filename, config)
	if err != nil {
		return config, err
	}	

	// Read user config haus yaml
	err = readCfg(usrcfgfile, config)
	if err != nil {
		return config, err
	}
	
	// Store the current path
	config.Pwd,err = os.Getwd()
	if err != nil {
		return config,err
	}
	return config, nil
}

// readCfg reads a file and parses it for yaml and unmarshals it into config.
func readCfg(cfgfile string, config *Config) (error) {
	// Read config from user home a overrite anything
	_,err := os.Stat(cfgfile)
	if err != nil {
	} else {
		cfg,err := ioutil.ReadFile(cfgfile)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(cfg, config)
		if err != nil {
			return err
		}
	}
	return nil
}