package haus

import(
	"testing"
	"os"
	"io/ioutil"
	"fmt"
)

func TestHausStruct( t *testing.T){
	testpath, err := ioutil.TempDir("", "haus")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)
	env := Environment{
		Name: "test",
		Variables: map[string]string{
			"test" : "testing",
		},
	}
	config := Config{
		Name: "test",
		Email: "test@test.com",
		Path: testpath,
		Environments: map[string]Environment{ "test": env},
	}
	config.Pwd = ".."
	fmt.Printf("%#v", config)
	haus := &Haus{ Config: config }
	err = haus.Run()
	if err != nil {
		t.Error(err)
	}

}