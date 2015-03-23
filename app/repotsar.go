package haus

import(
	"io/ioutil"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/SearchSpring/RepoTsar/gitutils"
)

type Signature struct {
	Name string
	Email string
}

type Repo struct {
	URL string
	Path string
	Branch string

}

func (r *Repo) CloneRepo(n string) error {
	fmt.Printf("Cloning repo: %s into %s\n",n,r.Path)
	cloneinfo := &gitutils.CloneInfo{
		Reponame: n,
		Path: r.Path,
		URL: r.URL,
		Branch: r.Branch,
	}
	_,err := cloneinfo.CloneRepo()
	if err != nil {
		return err
	}
	return nil
}

type RepoYml struct {
	Signature Signature
	Repos map[string]Repo
}

func (y *RepoYml) AddCfg(r map[string]Repo) {
	for k,v := range r {
		if y.Repos == nil {
			y.Repos = make(map[string]Repo)
		}
		y.Repos[k] = v
	}
}

// Write Yaml file
func (y *RepoYml) WriteYml(filename string) (string,error) {
	yaml,err := yaml.Marshal(y)
	if err != nil {
		return "",err
	}
	err = ioutil.WriteFile(filename, yaml, 0644)
	if err != nil {
		return "",err
	}
	return string(yaml[:]),err	
}


