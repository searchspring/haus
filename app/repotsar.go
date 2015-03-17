package haus

import(

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
	Yaml
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

