#!/bin/sh

sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' thirdparty/mcsql/Makefile
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' thirdparty/mrj/Makefile
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' sqlitedb/query.go
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' shm/worker.go
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' sqlitedb/sql_mapper.go

sudo make -C thirdparty/mcsql/
sudo make -C thirdparty/mrj/

sudo mkdir ${GOPATH}/src/elibot-apiserver
sudo cp -a * ${GOPATH}/src/elibot-apiserver

CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o elibot-server