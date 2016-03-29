FROM golang:alpine
ENV GO_VER go1.4.2.linux-amd64
ENV HAUS_VER 0.1.7
ENV GOPATH /app/go
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# Install build tools
RUN \
	 apk add --update \
		build-base \
		cmake \
		pkgconfig \
		# libssh2 \
		# libssh2-dev \
		openssl \
		openssl-dev \
		zlib \
		git \
		perl \
		bash \
		openssh-client \
		
		&& rm -rf /var/cache/apk/*


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
	# make test && \
	make install

ENV PKG_CONFIG_PATH=$GOPATH/src/gopkg.in/libgit2/git2go/vendor/libgit2/build/ 

# Install yaml,RepoTsar, and haus
RUN \
	go get gopkg.in/yaml.v2 && \ 
	go get github.com/SearchSpring/RepoTsar && \
	go get github.com/Searchspring/haus 

