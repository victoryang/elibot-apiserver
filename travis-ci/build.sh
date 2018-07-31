#!/bin/sh

export GOPATH=/root/go
export BUILDROOT=${GOPATH}/src/elibot-apiserver

# sed -i '/\"elibot-apiserver\/shm\"/d' main.go
# sed -i '/shms.Shutdown()/d' main.go
# sed -i '/shms := shm.NewServer(wss)/d' main.go
# sed -i '/shms.StartToWatch()/d' main.go
# sed -i 's/handleSignals(s, mcs, gs, wss, shms)/handleSignals(s, mcs, gs, wss)/g' main.go
# sed -i 's/, shms \*shm.ShmServer//g' main.go
# sed -i 's/shms,err := shm.NewServer(wss)/var err error = nil/g' main.go
# sed -i '/Log.Error(err.Error())/d' main.go

sudo cd ${BUILDROOT}
sudo cd thirdparty/mcsql/ && make
sudo cd ../mrj/ && make
sudo cd ../..
CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o elibot-server