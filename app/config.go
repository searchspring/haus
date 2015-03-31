package haus

import(
	"os"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents the configurations in haus config file.
type Config struct {
	Name string
	Email string
	Path string
	Pwd string
	Environments map[string]Environment
}

// ReadConfig reads the config file from the supplied full path and
// returns a Config and error.
func ReadConfig(filename string, usrcfgfile string )(*Config, error) {
	config := &Config{}
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(source, config)
	if err != nil {
		return config, err
	}

	// Read config from user home a overrite anything
	_,err = os.Stat(usrcfgfile)
	if err != nil {
	} else {
		usrcfg,err := ioutil.ReadFile(usrcfgfile)
		if err != nil {
			return config, err
		}
		err = yaml.Unmarshal(usrcfg, config)
		if err != nil {
			return config, err
		}
	}

	// Store the current path
	config.Pwd,err = os.Getwd()
	if err != nil {
		return config,err
	}
	
	return config, nil
}