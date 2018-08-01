# elibot-apiserver
[![Build Status](https://travis-ci.org/victoryang/elibot-apiserver.svg?branch=master)](https://travis-ci.org/victoryang/elibot-apiserver)

A api server for elibot


## docker
docker build -t apiserver:apiserver .
docker run -it --rm -v '${PWD}:/root/go/src/elibot-apiserver/' apiserver:apiserver