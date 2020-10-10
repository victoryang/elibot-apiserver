# elibot-apiserver
[![Build Status](https://travis-ci.org/victoryang/elibot-apiserver.svg?branch=master)](https://travis-ci.org/victoryang/elibot-apiserver)

A api server for elibot written in Go


## docker
	docker build -t apiserver:apiserver .
	chmod +x post-install.sh
	docker run -it --rm -v ${PWD}:/root/go/src/elibot-apiserver/ apiserver:apiserver

## build
	make

## Bump to go 1.12
    - check missing deps
       - export GO111MODULE=on
       - go list -m all
    - update deps
       - go mod edit -replace={old}[@v]={new}[@v]
    - update vendor
       - go mod vendor
