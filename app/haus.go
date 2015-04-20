// Package haus reads a yaml file describing docker architecture and
// generates from templates docker-compose and RepoTsar files.
package haus

import(
	"fmt"
	"os"
	"strings"
	"time"
)

// Haus represents an instance of the haus application.
type Haus struct{
	Config Config
}

type semaphore chan error

// Run starts the process of the haus application and returns error.
func (h *Haus) Run() error {
	c := h.Config
	
	repotsaryml := &RepoYml{
		Signature: Signature{
			Name: c.Name,
			Email: c.Email,
		},
	}
	dockeryml := &DockerYml{}
	// Create environments
	thrnum := len(c.Environments)
	sem := make(semaphore, thrnum)
	for k := range c.Environments {
		// Thread environment creatation
		go h.createEnv(k,dockeryml,repotsaryml,sem)
		// Sleep between thread creation so git doesn't step on it's self
		time.Sleep(50 * time.Millisecond)

	}
	// Wait for threads to finish
	for i := 0; i < thrnum; i++ { 
		err := <-sem
		if err != nil {
			return err
		}
	}

	// Create Yaml files
	_,err := repotsaryml.WriteYml(c.Path+"/repotsar.yml")
	if err != nil {
		return err
	}
	_,err = dockeryml.WriteYml(c.Path+"/docker-compose.yml")
	if err != nil {
		return err
	}
	return nil
		
}

// createEnv create a single environment define in the haus config.  Returns error.
func (h *Haus) createEnv(env string, dockeryml *DockerYml, repotsaryml *RepoYml, sem semaphore) error {
	// check requirements
	e := h.Config.Environments[env]
	for _,v := range e.Requirements {
		if _, ok := h.Config.Environments[v]; ok != true {
			err := fmt.Errorf("%#v requires %#v and it is not defined.",env,v)
			sem <-err
			return err
		}
	}
	name,version := nameSplit(env)
	// Create Cfgs from templates
	tmpl := &Template{
		Path: h.Config.Path,
		Pwd: h.Config.Pwd,
		Name: name,
		Branch: version,
		Version: version,
		Variables: e.Variables,
		Env: envMap(),
	}

	// Extract RepoTsar Cfgs from template
	repos,err := tmpl.RepoTsarCfgs()
	if err != nil {
		sem <-err
		return err
	}
	// Extract Docker Cfgs from template
	docker,err := tmpl.DockerCfgs()
	if err != nil {
		sem <-err
		return err
	}

	// Clone docker build repos first to ensure they are cloned into 
	// empty directories
	for dkey,dval := range docker {
		if dval.Build != "" {
			repo := repos[dkey]
			err = repo.CloneRepo(dkey)
			if err != nil {
				sem <-err
				return err
			}
			err = repo.CreateLink()
			if err != nil {
				sem <-err
				return err
			}
	
		} 
	}
	// Clone other repos after cloning docker build repos
	for rkey,rval := range repos {
		if _,ok := docker[rkey]; ok {
		} else {
			err = rval.CloneRepo(rkey)
			if err != nil {
				sem <-err
				return err
			}
			err = rval.CreateLink()
			if err != nil {
				sem <-err
				return err
			}			
		}
	}
	// Add Cfgs
	dockeryml.AddCfg(docker)
	repotsaryml.AddCfg(repos)
	sem <-nil
	return nil
}

// nameSplit splits string with _ into two strings if there's a '_'.
func nameSplit(n string) (string, string) {
		// split name and version or branch
		s := strings.Split(n,"_")
		var version string
		var name string
		if len(s) > 1 {
			name, version = s[0], s[1]
		} else {
			name = s[0]
		}
		return name, version
}

// envMap returns a map[string]string of Environment variables
func envMap() map[string]string {
	env := make(map[string]string)
	envstrings := os.Environ()
	for i := range envstrings {
		kv := strings.Split(envstrings[i],"=")
		env[kv[0]] = kv[1]
	}
	return env
} 