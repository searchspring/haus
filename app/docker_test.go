package haus

import(
	"testing"
	"os"
	"io/ioutil"
	"regexp"
	// "fmt"
)

func TestDockYmlStruct(t *testing.T) {
	dockeryml := &DockerYml{}
	dockercfg := map[string]DockerCfg{
		"test2": DockerCfg{
			Build: "/test",
			Dns: []string{
				"8.8.8.8",
			},
			Ports: []string{
				"80:80",
			},
			Volumes: []string{
				"test:test",
			},
			Expose: []string{
				"80",
			},
		},
	}
	// AddCfg
	dockeryml.AddCfg(dockercfg)

	// WriteYml
	testpath, err := ioutil.TempDir("", "haus")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)
	_,err = dockeryml.WriteYml(testpath+"/test.yml")
	if err != nil {
		t.Error(err)
	}
	yamlfile,err := ioutil.ReadFile(testpath+"/test.yml")
	if err != nil {
		t.Error(err)
	}
	teststring := "test2:\n  build: /test\n  dns:\n  - 8.8.8.8\n  expose:\n  - \"80\"\n  ports:\n  - 80:80\n  volumes:\n  - test:test\n"
	match,err := regexp.Match(teststring,yamlfile)
	if err != nil {
		t.Error(err)
	}
	if match != true {
		t.Error("Yamlfile missing contents")
	}
}