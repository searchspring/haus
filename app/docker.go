package haus

import(
)

type DockerCfg struct {
	Image string `yaml:"image,omitempty"`
	Build string `yaml:"build,omitempty"`
	Command string `yaml:"build,omitempty"`
	Links []string `yaml:"build,omitempty"`
	External_links []string `yaml:"build,omitempty"`
	Dns []string `yaml:"dns,omitempty"`
	Dns_server []string `yaml:"build,omitempty"`
	Ports []string `yaml:"ports,omitempty"`
	Volumes []string `yaml:"volumes,omitempty"`
	Volumes_from []string `yaml:"build,omitempty"`
	Env_file []string `yaml:"build,omitempty"`
	Net string `yaml:"build,omitempty"`
	Cap_add []string `yaml:"build,omitempty"`
	Cap_drop []string `yaml:"build,omitempty"`
	Expose []string `yaml:"expose,omitempty"`
	Environment []string `yaml:"environment,omitempty"`
	Working_dir string `yaml:"build,omitempty"`
	Entrypoint string `yaml:"build,omitempty"`
	User string `yaml:"build,omitempty"`
	Hostname string `yaml:"build,omitempty"`
	Domainname string `yaml:"build,omitempty"`
	Mem_limit string `yaml:"build,omitempty"`
	Privileged string `yaml:"build,omitempty"`
	Restart string `yaml:"build,omitempty"`
	Stdin_open string `yaml:"build,omitempty"`
	Tty string `yaml:"build,omitempty"`
	Cpu_shares string `yaml:"build,omitempty"`
}

type DockerYml struct {
	Yaml
	Cfg map[string]DockerCfg
}

func (y *DockerYml) AddCfg(d map[string]DockerCfg) {
	for k,v := range d {
		if y.Cfg == nil {
			y.Cfg = make(map[string]DockerCfg)
		}
		y.Cfg[k] = v
	}
}

func (y *DockerYml) Cfgs() map[string]DockerCfg {
	return y.Cfg
}
