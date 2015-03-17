package haus

import(
	"testing"
	"os"
	"io/ioutil"
)

func TestDockYmlStruct(t *testing.T) {
	dockeryml := &DockerYml{}
	dockercfg := map[string]DockerCfg{
		"test": DockerCfg{
			Build: "./",
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

}