haus
========

Keep your haus in order.  Haus orchestrates local docker development environments.  With a simple Yaml config, and a series of templates, haus will dynamically build RepoTsar and Docker-Compose config files for you. 

Installing
==========

There are static binary packages for OSX in releases.  If you would like to build haus yourself, follow the instructions below.

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

#### Building libgit2 and git2go

```bash
go get -d gopkg.in/libgit2/git2go.v22
cd $GOPATH/src/gopkg.in/libgit2/git2go.v22
git submodule update --init 
make install
```

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

# haus.yml
Haus will look in the current directory for a haus.yml file.  This file defines the architecture you want to build.  Here is an example:

```yaml
name: "Your Name"
email: email@address.com

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
      - 172.17.42.1
      - 8.8.8.8
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
        - 172.17.42.1
        - 8.8.8.8
    volumes:
        - {{.Path}}/src/docker-node/src/ghost/:/usr/src/ghost/
    environment:
        - SERVICE_2368_NAME=ghost
        - SERVICE_TAG=ghost
    
    net: "bridge"
    command: "npm install --production&& npm start --production"
```

# User Config

Haus looks in ~/.hauscfg.yml for user configuration.  This a YAML file that should look like the following.

```YAML
name: "Your Name"
email: email@domain.com
hausrepo: git@github.com:SearchSpring/haus.git
```

The name and email settings are used as your git signature.  The hausrepo setting should be the git address of a repo where you have a haus.yml and templates for haus.  If you use the haus command in a empty directory the repo specified in hausrepo will be cloned into the current directory and then haus will use that haus.yml config.  You may also specify -branch with the haus command to checkout a specific branch of the repo you configured hausrepo for.


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
