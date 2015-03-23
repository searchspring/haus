package haus

import(
	"testing"
)

func TestTemplateStruct(t *testing.T){

	tmpl := &Template{
		Path: "/tmp/test",
		Pwd: "..",
		Name: "test",
		Branch: "master",
		Version: "1.1.1",
		Variables: map[string]string{
			"Clustername": "test",
		},
	}
	_,err := tmpl.Parse()
	if err != nil {
		t.Error(err)
	}

	_,err = tmpl.DockerCfgs()
	if err != nil {
		t.Error(err)
	}
	_,err = tmpl.RepoTsarCfgs()
	if err != nil {
		t.Error(err)
	}
}