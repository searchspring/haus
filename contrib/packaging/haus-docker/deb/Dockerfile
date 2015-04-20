FROM buildpack-deps:jessie-scm

ENV GO_VER go1.4.2.linux-amd64
ENV HAUS_VER 0.1.4
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
		zlibc 


# Install git2go
RUN \
	go get -d gopkg.in/libgit2/git2go.v22 && \
	cd $GOPATH/src/gopkg.in/libgit2/git2go.v22 && \
	git checkout next && \
	git submodule update --init && \
	make install

RUN \
	go get github.com/Searchspring/haus && \
	go get github.com/Searchspring/RepoTsar

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
