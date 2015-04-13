package main

import(
	"log"
	"fmt"
	"flag"
	"path/filepath"

	"github.com/SearchSpring/haus/fileutils"
	"github.com/SearchSpring/haus/app"
)

var configfile string
var path string
var branch string
var version bool
const (
	versioninfo = "v0.1.2"
)

func main(){
	flag.StringVar(&configfile, "config", "haus.yml", "YAML config file")
	flag.StringVar(&path, "path", "./hauscfg", "Path to generate files in")
	flag.StringVar(&branch, "branch", "master", "git branch for hausrepo")
	flag.BoolVar(&version, "version",false,"haus version")
	flag.Parse()

	if version == true {
		fmt.Printf("haus version %s\n", versioninfo)
		return
	}

	config,err := haus.ReadConfig(configfile,"~/.hauscfg.yml", branch)
	if err != nil {
		log.Fatalf("\n%s",err)
	}
	
	config.Path,err = filepath.Abs(path)
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

