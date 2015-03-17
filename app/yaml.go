package haus

import(
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Yaml struct{
}

func (y *Yaml) Cfgs() Yaml {
	return *y
}

// Write Yaml file
func (y *Yaml ) WriteYml(filename string) (string,error) {
	yaml,err := yaml.Marshal(y.Cfgs())
	if err != nil {
		return "",err
	}
	err = ioutil.WriteFile(filename, yaml, 0644)
	if err != nil {
		return "",err
	}
	return string(yaml[:]),err	
}
