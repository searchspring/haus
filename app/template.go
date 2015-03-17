package haus

import(
	"bytes"
	"io/ioutil"
	"text/template"
	
	"gopkg.in/yaml.v2"
)

type Template struct {
	Path string
	Image string
	Name string
	Branch string
	Version string
	Variables map[string]string
	Parsed ParsedTmpl
	

}

type ParsedTmpl struct {
	Repotsar map[string]Repo
	Docker map[string]DockerCfg
}

func (t *Template) Parse() (ParsedTmpl, error){
	if &t.Parsed != nil {
	 	return t.Parsed, nil
	}
	source, err := ioutil.ReadFile("./templates/"+t.Name+".yml.tmpl")
	if err != nil {
		return t.Parsed, err
	}
	template,err := template.New(t.Name).Parse(string(source[:]))
	if err != nil {
		return t.Parsed, err
	}
	buf := &bytes.Buffer{}
	err = template.Execute(buf,t)
	if err != nil {
		return t.Parsed,err
	}
	parsed := ParsedTmpl{}
	err = yaml.Unmarshal(buf.Bytes(),parsed)
	if err != nil {
		return t.Parsed,err
	}
	t.Parsed = parsed
	return parsed, nil

}

func (t *Template) DockerCfgs() (map[string]DockerCfg, error) {
	if &t.Parsed == nil {
		_,err := t.Parse()
		if err != nil {
			return nil,err
		}
	}
	return t.Parsed.Docker,nil 
}

func (t *Template) RepoTsarCfgs() (map[string]Repo, error) {
	if &t.Parsed == nil {
		_,err := t.Parse()
		if err != nil {
			return nil,err
		}
	}
	return t.Parsed.Repotsar,nil
}

