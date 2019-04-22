# spa(go + vue.js)

## Installation
### Go Environment
```bash
    GO_VERSION=1.12
    GO_DOWNLOAD_URL=https://dl.google.com/go

    GOROOT=/usr/local/go
    GOPATH=/root/go
    PATH=${GOPATH}/bin:${GOROOT}/bin:${PATH}

    curl -s ${GO_DOWNLOAD_URL}/go${GO_VERSION}.linux-amd64.tar.gz | tar -v -C /usr/local/ -xz
```

### Node Environment
```bash
    VERSION=v10.15.0
    DISTRO=linux-x64
    PATH=/usr/local/lib/nodejs/node-$VERSION-$DISTRO/bin:$PATH

    wget https://nodejs.org/dist/${VERSION}/node-${VERSION}-${DISTRO}.tar.xz
    mkdir -p /usr/local/lib/nodejs
    tar -xJvf node-$VERSION-$DISTRO.tar.xz -C /usr/local/lib/nodejs
```