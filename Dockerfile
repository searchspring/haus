FROM golang:1.6.0
ENV GO_VER go1.4.2.linux-amd64
ENV HAUS_VER 0.1.7
ENV GOPATH /app/go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# Install build tools
RUN \
	apt-get update && \
	apt-get -y install \
		make \
		cmake \
		pkg-config \
		libssh2-1-dev \
		libssh2-1 \
		libssl-dev \
		openssh-client \
		openssh-server \
		zlibc \
		git


# Install git2go and libgit2
RUN 	go get -d gopkg.in/libgit2/git2go.v22 && \
	cd $GOPATH/src/gopkg.in/libgit2/git2go.v22 && \
	git checkout next && \
	git submodule update --init && \
	./script/build-libgit2-static.sh && \ 
	cd vendor/libgit2/build/ && \
	make && \
	make install && \
	cd ../../../ && \
	make test && \
	make install

ENV PKG_CONFIG_PATH=$GOPATH/src/gopkg.in/libgit2/git2go/vendor/libgit2/build/ 

# Install yaml,RepoTsar, and haus
RUN \
	go get gopkg.in/yaml.v2 && \ 
	go get github.com/SearchSpring/RepoTsar && \
	mkdir ${GOPATH}/src/github.com/SearchSpring/haus

COPY . ${GOPATH}/src/github.com/SearchSpring/haus/

RUN \
	cd ${GOPATH}/src/github.com/SearchSpring/haus/ && \
	go test ./... && \
	go install 

WORKDIR /mnt
COPY start /bin/start

ENV HAUSPATH /var/tmp
ENTRYPOINT [ "/bin/start" ]