package haus

import(
	"io/ioutil"
	"os"
	"testing"
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
	}
	err = repo.CloneRepo("test")
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
	}
	repoyml := &RepoYml{}
	repocfg := map[string]Repo{ "test": repo}
	repoyml.AddCfg(repocfg)
	_,err = repoyml.WriteYml(testpath+"/test.yml")
	if err != nil {
		t.Error(err)
	}

}