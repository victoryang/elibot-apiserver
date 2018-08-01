# elibot-apiserver
[![Build Status](https://travis-ci.org/victoryang/elibot-apiserver.svg?branch=master)](https://travis-ci.org/victoryang/elibot-apiserver)

A api server for elibot


## docker
docker build -t apiserver:apiserver .
chmod +x post-install.sh
docker run -it --rm -v '${PWD}:/root/go/src/elibot-apiserver/' apiserver:apiserver

## build
chmod +x build/build.sh & sh -C build/build.sh