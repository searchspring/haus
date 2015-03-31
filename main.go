package main

import(
	"log"
	"flag"
	"path/filepath"

	"github.com/SearchSpring/haus/fileutils"
	"github.com/SearchSpring/haus/app"
)

var configfile string
var path string

func main(){
	flag.StringVar(&configfile, "config", "haus.yml", "YAML config file")
	flag.StringVar(&path, "path", "./hauscfg", "Path to generate files in")
	config,err := haus.ReadConfig(configfile,"~/.hauscfg.yml")
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

