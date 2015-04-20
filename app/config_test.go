package haus

import (
	"os"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestReadConfig(t *testing.T) {

	// Test missing config
	_,err := ReadConfig("bogusbogus","","master", "./hauscfg", make(map[string]string) )
	if err == nil {
		t.Error("Expected missing file error, didn't get it.")
	}

	// Test broken config
	_,err = ReadConfig("config.go","", "master", "./hauscfg", make(map[string]string) )
	if err == nil {
		t.Error("Expected yaml error, didn't get it.")
	}

	// Test normal config
	config,err := ReadConfig("../haus.yml","", "master", "./hauscfg", make(map[string]string) )
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
	testpath_usrcfg, err := ioutil.TempDir("", "haus_usrcfg")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath_usrcfg)

	usrCfg := map[string]string{
		"name": "Testy Testerson",
		"email": "Test@test.com",
		"path": "/my/test/path",
		"hausrepo": "https://git@github.com/SearchSpring/haus.git",
	}
	yaml,err := yaml.Marshal(usrCfg)
	if err != nil {
		t.Error(err)
	}
	usrcfgfile := testpath_usrcfg+"/.hauscfg.yml"
	err = ioutil.WriteFile(usrcfgfile, yaml, 0644)
	if err != nil {
		t.Error(err)
	}
	
	// Test autocheckout
	testpath, err := ioutil.TempDir("", "haus")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)
	pwd,err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	err = os.Chdir(testpath)
	if err != nil {
		t.Error(err)
	}
	defer os.Chdir(pwd)

	config,err = ReadConfig("bogus", usrcfgfile, "master", "./hauscfg", make(map[string]string))
	if err != nil {
		t.Error(err)
	} 

	// Test with config	
	if config.Name != "Testy Testerson" {
		t.Error("User config failed, expected Testy Testerson, got ",config.Name)
	}

	config,err = ReadConfig("./haus.yml", usrcfgfile, "master", "./hauscfg", make(map[string]string) )
	if err != nil {
		t.Error(err)
	}

	if config.Name != "Testy Testerson" {
		t.Error("User config failed, expected Testy Testerson, got ",config.Name)
	}
}