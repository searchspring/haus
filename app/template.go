package haus

import(
	"bytes"
	"io/ioutil"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Template represents a single instance of a haus template file.
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

// ParsedTmpl represents a collection of Repo and DockerCfgs created from haus templates.
type ParsedTmpl struct {
	Repotsar map[string]Repo
	Docker map[string]DockerCfg
}

// Parse processes template files to create Repo and DockerCfgs, returns a ParsedTmpl.
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

// DockerCfgs returns DockerCfgs stored in Template.
func (t *Template) DockerCfgs() (map[string]DockerCfg, error) {
	parsed,err := t.Parse()
	return parsed.Docker,err 
}

// RepoTsarCfgs returns RepoTsarCfgs stored in Template.
func (t *Template) RepoTsarCfgs() (map[string]Repo, error) {
	parsed,err := t.Parse()
	return parsed.Repotsar,err
}

