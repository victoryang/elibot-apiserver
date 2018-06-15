#!/bin/sh

sed -i '/\"elibot-apiserver\/shm\"/d' main.go
sed -i '/shms.Shutdown()/d' main.go
sed -i '/shms := shm.NewServer(wss)/d' main.go
sed -i '/shms.StartToWatch()/d' main.go
sed -i 's/handleSignals(s, mcs, gs, wss, shms)/handleSignals(s, mcs, gs, wss)/g' main.go
sed -i 's/, shms \*shm.ShmServer//g' main.go
sed -i 's/shms,err := shm.NewServer(wss)/var err error = nil/g' main.go
sed -i '/Log.Error(err.Error())/d' main.go

go build -o elibot-server