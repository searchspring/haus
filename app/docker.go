package haus

import(
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type DockerCfg struct {
	Build string `yaml:"build,omitempty"`
	Cap_add []string `yaml:"cap_add,omitempty"`
	Cap_drop []string `yaml:"cap_drop,omitempty"`
	Command string `yaml:"command,omitempty"`
	Cpu_shares string `yaml:"cpu_shares,omitempty"`
	Dns []string `yaml:"dns,omitempty"`
	Dns_search []string `yaml:"dns_search,omitempty"`
	Domainname string `yaml:"domainname,omitempty"`
	Entrypoint []string `yaml:"entrypoint,omitempty"`
	Env_file []string `yaml:"env_file,omitempty"`
	Environment []string `yaml:"environment,omitempty"`
	Expose []string `yaml:"expose,omitempty"`
	External_links []string `yaml:"external_links,omitempty"`
	Hostname string `yaml:"hostname,omitempty"`
	Image string `yaml:"image,omitempty"`
	Links []string `yaml:"links,omitempty"`
	Mem_limit string `yaml:"mem_limit,omitempty"`
	Net string `yaml:"net,omitempty"`
	Ports []string `yaml:"ports,omitempty"`
	Privileged string `yaml:"privileged,omitempty"`
	Restart string `yaml:"restart,omitempty"`
	Stdin_open string `yaml:"stdin_open,omitempty"`
	Tty string `yaml:"tty,omitempty"`
	User string `yaml:"user,omitempty"`
	Volumes []string `yaml:"volumes,omitempty"`
	Volumes_from []string `yaml:"volumes_from,omitempty"`
	Working_dir string `yaml:"working_dir,omitempty"`
}

type DockerYml struct {
	Cfg map[string]DockerCfg
}

func (y *DockerYml) AddCfg(d map[string]DockerCfg) {
	for k,v := range d {
		if y.Cfg == nil {
			y.Cfg = d 
		} else {
			y.Cfg[k] = v
		}
	}
}

func (y *DockerYml) Cfgs() map[string]DockerCfg {
	return y.Cfg
}

// Write Yaml file
func (y *DockerYml) WriteYml(filename string) (string,error) {
	yaml,err := yaml.Marshal(y.Cfg)
	if err != nil {
		return "",err
	}
	err = ioutil.WriteFile(filename, yaml, 0644)
	if err != nil {
		return "",err
	}
	return string(yaml[:]),err	
}
