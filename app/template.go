package haus

import(
	"bytes"
	"io/ioutil"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Template struct {
	Path string
	Pwd string
	Image string
	Name string
	Branch string
	Version string
	Variables map[string]string
	Parsed *ParsedTmpl
	

}

type ParsedTmpl struct {
	Repotsar map[string]Repo
	Docker map[string]DockerCfg
}

func (t *Template) Parse() (ParsedTmpl, error){
	if t.Parsed != nil {
		return *t.Parsed,nil
	}
	// Read template
	templatefile := t.Pwd+"/templates/"+t.Name+".yml.tmpl"
	source, err := ioutil.ReadFile(templatefile)
	if err != nil {
		return ParsedTmpl{}, err
	}

	// Parse template
	template,err := template.New(t.Name).Parse(string(source[:]))
	if err != nil {
		return ParsedTmpl{}, err
	}
	// 
	buf := &bytes.Buffer{}
	err = template.Execute(buf,t)
	if err != nil {
		return ParsedTmpl{},err
	}
	parsed := &ParsedTmpl{}
	err = yaml.Unmarshal(buf.Bytes(),parsed)
	if err != nil {
		return *parsed,err
	}
	t.Parsed = parsed
	return *parsed, nil

}

func (t *Template) DockerCfgs() (map[string]DockerCfg, error) {
	parsed,err := t.Parse()
	return parsed.Docker,err 
}

func (t *Template) RepoTsarCfgs() (map[string]Repo, error) {
	parsed,err := t.Parse()
	return parsed.Repotsar,err
}

