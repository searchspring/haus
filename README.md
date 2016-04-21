haus
========

Keep your haus in order.  Haus gives you an easy way to recreate your docker development environments from scratch and share them with others.  With a simple Yaml config, and a series of templates, haus will dynamically build RepoTsar and Docker-Compose config files for you. 

Docker
======
Instead of building or installing haus, it's simplest to run haus via docker.  Follow instructions here https://hub.docker.com/r/searchspring/haus/ 

Installing
==========

There are static binary packages for OSX and Debian in releases.  If you would like to build haus yourself, follow the instructions below.

# Requirements
This project is written in go.

### Go

https://golang.org/doc/install

Don't forget to set your GOPATH and add the locations of the go bin to your PATH.  You should also ensure that your PKG_CONFIG_PATH is set for compiling purposes later on.  Something like this in your .bashrc :

```bash
export GOPATH=$HOME/code/go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin/
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:/usr/lib/pkgconfig
```

### Docker

If you're on OSX or Windows checkout boot2docker http://boot2docker.io/

### Installing git2go

This project requires git2go.v22, which in turn requires libgit2.  If you want SSH support with this application, ensure you have libssh2 installed as well.  A recent version of cmake is required to build libgit2 http://www.cmake.org/download/ .  You'll need to commandline tools for cmake, if you're on OSX and you've installed the GUI, goto Tools->Install for Command Line Use.  This may throw an error about failing to create symlinks.  This is a permissions error, you can create this links by hand.

```bash
/usr/bin/
cmake -> /Applications/CMake.app/Contents/bin/cmake
cmake-gui -> /Applications/CMake.app/Contents/bin/cmake-gui
cmakexbuild -> /Applications/CMake.app/Contents/bin/cmakexbuild
```

Follow instructions at https://github.com/libgit2/git2go to build libgit2 and git2go.

#### Installing yaml

This project also requires yaml.v2

```bash
go get gopkg.in/yaml.v2
```

#### Installing RepoTsar

```bash
go get github.com/SearchSpring/RepoTsar
``` 

Usage
=====

# ~/.hauscfg.yml

If haus is run in an empty directory it will pull the config from ~/.hauscfg.yml.  The following is an example of what that file should look like.

```yaml
name: "Your Name"
email: email@address.com
hausrepo: git@bitbucket.org:yourrepo/haus-yaml.git
```

The **hausrepo** option should be configured with a repo containing haus config files.  Haus will **one time** checkout this repo and then run haus.  Anytime a haus.yml file exists in the current directory, this setting is ignored.  If you pass *-branch* **branchname** as an argument, haus will checkout that specific branch.


# haus.yml
Haus will look in the current directory for a haus.yml file.  This file defines the architecture you want to build.  Here is an example:

```yaml

# Git Signature Info
name: "Your Name"
email: email@address.com

# Global Defaults
variables:
  dns1: 172.17.42.1
  dns2: 8.8.8.8

# Environment Definitions
environments:
  mysql_latest:
  redis_latest:
  elasticsearch_1.3.4:
    variables:
      name: "ESTest1"
      clustername: "test-cluster1"
      port: 9201
  elasticsearch_1.1.1:
    variables:
      name: "ESTest2"
      clustername: "test-cluster2"
      port: 9202
  consul:
  registrator:
    requirements:
        - consul
```

The 'name' and 'email' fields will be used in your git signature for any checkouts that are required.  Each element under 'environments' corresponds to a template in a templates directory in the current path.  The template file name will match the element name up to an `_`, plus .yml.tmpl as a suffix.  Example, in the above config, the template file for elasticsearch_1.3.4 will be
`./templates/elasticsearch.yml.tmpl`
If there is an `_` in the name element, it will be passed to the template as the variable 'Version'.  Any additional variables you would like to pass to the template can be defined under 'Variables' and 'Branch'.  Requirements is completely optional, it is an array of names, requiring that those definitions exist in this document.  This allows you to show a relationship between different definitions.  In the example above, the registrator definition requires that a consul definition exists.

# Templates

Templates live in a directory `templates` in your current path.  Each template defines either a git repository for RepoTsar to manage and/or a docker instance.

Here is an example of a docker only template
```yaml
docker:
  {{.Variables.name}}ES:
    image: searchspring/elasticsearch:consul-config-{{.Version}}
    dns:
      - {{.Variables.dns1}}
      - {{.Variables.dns2}}
    ports:
      - "{{.Variables.port}}:9200"
    expose:
      - "9200"
    environment:
      - SERVICE_9200_NAME=es
      - CLUSTERNAME={{.Variables.clustername}}
```


Here is an example of a template with a repo for the dockers instance and for an application.  This is generate a docker-compose.yml that will build the docker image from the repo and then mount a volume with the application code from the application repo.

```yaml
repotsar: 
  nodejs:
    url: ssh://git@github.com/joyent/docker-node.git
    path: {{.Path}}/src/docker-node/
    branch: {{.Branch}}
  ghost:
    url: ssh://git@github.com/TryGhost/Ghost.git
    path: {{.Path}}/src/docker-node/src/ghost
    branch: {{.Branch}}
docker:
  test:
    build: {{.Path}}/src/docker-node/
    ports:
        - "80:2368"
    expose:
        - "2368"
    dns:
        - {{.Variables.dns1}}
        - {{.Variables.dns2}}
    volumes:
        - {{.Path}}/src/docker-node/src/ghost/:/usr/src/ghost/
    environment:
        - SERVICE_2368_NAME=ghost
        - SERVICE_TAG=ghost
    
    net: "bridge"
    command: "npm install --production&& npm start --production"
```

Inside the templates the following variables may be used

* .Path - This will be the value supplied to the command line arguement -path (default ./hauscfg)
* .Version / .Branch - This will be any value after the first `_` in the name of the environment
* .Variables - Any variable defined under the `variables` section of the enviroment definition
* .Env - Any shell environment variables 


# User Config

Haus looks in ~/.hauscfg.yml for user configuration.  This a YAML file that should look like the following.  The variables section in this file will override the global defaults section in the haus.yml.

```YAML
name: "Your Name"
email: email@domain.com
hausrepo: git@github.com:SearchSpring/haus.git

variables:
  dns1: 172.17.42.1
  dns2: 8.8.8.8

```

The name and email settings are used as your git signature.  The hausrepo setting should be the git address of a repo where you have a haus.yml and templates for haus.  If you use the haus command in a empty directory the repo specified in hausrepo will be cloned into the current directory and then haus will use that haus.yml config.  You may also specify -branch with the haus command to checkout a specific branch of the repo you configured hausrepo for.

# Syntax

## Usage of haus:
```
  -branch="master": git branch for hausrepo
  -config="haus.yml": YAML config file
  -path="./hauscfg": Path to generate files in
  -variables=map[]: variables passed to templates
  -version=false: haus version
```

### branch
  This option allows you to specify what branch haus should use when cloning from the hausrepo specified in your ~/.hauscfg.yml

### config
  This option allows you to tell haus to use a haus.yml file in a different location other than current directory.

### path
  This option specifies a different location for haus config files to be created in rather than ./hauscfg.

### variables
  Using this option allows you to override variables in the templates from the command line.  Example:
  ```bash
  bash-3.2$ haus -variables dns1=1.1.1.1 -variables dns2=2.2.2.2
  ```
### version
  Prints the version of haus and exits.

License and Author
==================

* Author:: Greg Hellings (<greg@thesub.net>)


Copyright 2014, B7 Interactive, LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
