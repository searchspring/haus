package haus

import(
)

type DockerCfg struct {
	Image string `yaml:"image,omitempty"`
	Build string `yaml:"build,omitempty"`
	Dns []string `yaml:"dns,omitempty"`
	Ports []string `yaml:"ports,omitempty"`
	Volumes []string `yaml:"volumes,omitempty"`
	Expose []string `yaml:"expose,omitempty"`
	Environment []string `yaml:"environment,omitempty"`

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
