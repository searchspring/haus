FROM buildpack-deps:jessie-scm

ENV GO_VER go1.4.2.linux-amd64
ENV HAUS_VER 0.1.6
ENV GOPATH /app/go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# Download go
RUN \
	cd /tmp && \
	wget https://storage.googleapis.com/golang/${GO_VER}.tar.gz

# install go
RUN \
	mkdir -p $GOROOT && \
	mkdir -p $GOPATH/bin && \
	mkdir -p $GOPATH/src && \
	cd /usr/local && \
	tar zxvf /tmp/${GO_VER}.tar.gz && \
	rm /tmp/${GO_VER}.tar.gz

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
		git \
		libcurl4-openssl-dev


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
	go get github.com/Searchspring/haus 

# Build Package
RUN \
	mkdir -p /app/deb && \
	cd /app/deb/ && \
	mkdir -p haus_${HAUS_VER}/usr/local/bin && \
	cp $GOPATH/bin/haus haus_${HAUS_VER}/usr/local/bin/ && \
	mkdir -p haus_${HAUS_VER}/DEBIAN

COPY control /app/deb/haus_${HAUS_VER}/DEBIAN/

RUN \
	cd /app/deb && \
	dpkg-deb --build haus_${HAUS_VER}

COPY .hauscfg.yml /root/
