#!/bin/sh

# sed -i '/\"elibot-apiserver\/shm\"/d' main.go
# sed -i '/shms.Shutdown()/d' main.go
# sed -i '/shms := shm.NewServer(wss)/d' main.go
# sed -i '/shms.StartToWatch()/d' main.go
# sed -i 's/handleSignals(s, mcs, gs, wss, shms)/handleSignals(s, mcs, gs, wss)/g' main.go
# sed -i 's/, shms \*shm.ShmServer//g' main.go
# sed -i 's/shms,err := shm.NewServer(wss)/var err error = nil/g' main.go
# sed -i '/Log.Error(err.Error())/d' main.go

sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' thirdparty/mcsql/Makefile
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' thirdparty/mrj/Makefile
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' sqlitedb/query.go
sed -i 's/\/root\/mcserver/\/usr\/local\/include\/mcserver/g' shm/worker.go

sudo make -C thirdparty/mcsql/
sudo make -C thirdparty/mrj/

sudo mkdir ${GOPATH}/src/elibot-apiserver
sudo cp -a * ${GOPATH}/src/elibot-apiserver

CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o elibot-server