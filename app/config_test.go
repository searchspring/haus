package haus

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	_,err := ReadConfig("bogusbogus")
	if err == nil {
		t.Error("Expected missing file error, didn't get it.")
	}
	_,err = ReadConfig("config.go")
	if err == nil {
		t.Error("Expected yaml error, didn't get it.")
	}
	config,err := ReadConfig("../haus.yml")
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
	
}