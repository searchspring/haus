package haus

import (
	"os"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestReadConfig(t *testing.T) {

	// Test missing config
	_,err := ReadConfig("bogusbogus","")
	if err == nil {
		t.Error("Expected missing file error, didn't get it.")
	}

	// Test broken config
	_,err = ReadConfig("config.go","")
	if err == nil {
		t.Error("Expected yaml error, didn't get it.")
	}

	// Test normal config
	config,err := ReadConfig("../haus.yml","")
	if err != nil {
		t.Error(err)
	}
	v := config.Name
	if v != "Your Name" {
		t.Error("Expected Your Name, got ", v)
	}
	v = config.Email
	if v != "email@address.com" {
		t.Error("Expected email@address.com, got ",v)
	}

	// Test user config with normal config
	testpath, err := ioutil.TempDir("", "haus")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)

	usrCfg := map[string]string{
		"name": "Testy Testerson",
		"email": "Test@test.com",
		"path": "/my/test/path",
	}
	yaml,err := yaml.Marshal(usrCfg)
	if err != nil {
		t.Error(err)
	}
	usrcfgfile := testpath+"/.hauscfg.yml"
	err = ioutil.WriteFile(usrcfgfile, yaml, 0644)
	if err != nil {
		t.Error(err)
	}

	config,err = ReadConfig("../haus.yml", usrcfgfile)
	if err != nil {
		t.Error(err)
	}

	if config.Name != "Testy Testerson" {
		t.Error("User config failed, expected Testy Testerson, got ",config.Name)
	}


	
}