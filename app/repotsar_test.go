package haus

import(
	"io/ioutil"
	"os"
	"testing"
	"regexp"
)

func TestRepoStruct(t *testing.T) {
	testpath, err := ioutil.TempDir("", "haus")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)

	repo := Repo{
		Path: testpath+"/testrepo",
		URL: "ssh://git@github.com/libgit2/git2go.git",
		Branch: "master",
		Link: testpath+"/testrepo-link",
	}
	err = repo.CloneRepo("test")
	if err != nil {
		t.Error(err)
	}
	err = repo.CreateLink()
	if err != nil {
		t.Error(err)
	}
	_,err = os.Lstat(testpath+"/testrepo-link")
	if err != nil {
		t.Error(err)
	}
}

func TestRepoYmlStruct(t *testing.T){
	testpath, err := ioutil.TempDir("", "haus")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(testpath)
	repo := Repo{
		Path: testpath+"/testrepo",
		URL: "ssh://git@github.com/libgit2/git2go.git",
		Branch: "master",
		Link: testpath+"/testrepo-link",
	}
	repoyml := &RepoYml{}
	repocfg := map[string]Repo{ "test": repo}
	repoyml.AddCfg(repocfg)

	// Write Yaml
	_,err = repoyml.WriteYml(testpath+"/test.yml")
	if err != nil {
		t.Error(err)
	}
	yamlfile,err := ioutil.ReadFile(testpath+"/test.yml")
	if err != nil {
		t.Error(err)
	}
	match,err := regexp.Match(".*github.*",yamlfile)
	if err != nil {
		t.Error(err)
	}
	if match != true {
		t.Error("Yamlfile missing contents")
	}
}