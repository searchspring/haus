package main

import(
	"log"
	"fmt"
	"flag"
	"strings"
	"path/filepath"

	"github.com/SearchSpring/haus/fileutils"
	"github.com/SearchSpring/haus/app"
)

var configfile string
var path string
var branch string
var version bool
const (
	versioninfo = "v0.1.4"
)

type vars map[string]string
var variables vars

func ( v *vars ) String() string {
	return fmt.Sprintf("%s", *v)
}

func ( v *vars ) Set(value string) error {
	keyval := strings.Split(value, "=")
	(*v)[keyval[0]] = keyval[1]
	return nil
}

func main(){
	variables = make(map[string]string)
	flag.StringVar(&configfile, "config", "haus.yml", "YAML config file")
	flag.StringVar(&path, "path", "./hauscfg", "Path to generate files in")
	flag.StringVar(&branch, "branch", "master", "git branch for hausrepo")
	flag.Var(&variables, "variables", "variables passed to templates")
	flag.BoolVar(&version, "version",false,"haus version")
	flag.Parse()

	if version == true {
		fmt.Printf("haus version %s\n", versioninfo)
		return
	}

	abspath,err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("\n%s",err)
	}

	config,err := haus.ReadConfig(configfile,"~/.hauscfg.yml", branch, abspath, variables)
	if err != nil {
		log.Fatalf("\n%s",err)
	}

	_,err = fileutils.CreatePath(config.Path)
	if err != nil {
		log.Fatalf("\n%s",err)
	}
	
	haus := haus.Haus{
		Config: *config,
	}
	err = haus.Run()
	if err != nil {
		log.Fatalf("\n%s",err)
	}
}

