#!/bin/sh

BASE=elibot-server
API=api
API_V1=${BASE}/${API}/v1
API_V2=${BASE}/${API}/v2
API_V3=${BASE}/${API}/v3
CONFIG=${BASE}/config
DB=${BASE}/db
LOG=${BASE}/log
MCSERVER=${BASE}/mcserver
MIDDLEWARE=${BASE}/middleware
SERVERPB=${BASE}/serverpb
SQLITEDB=${BASE}/sqlitedb
VENDOR=${BASE}/vendor

sed -i '/\"elibot-apiserver\/shm\"/d' main.go
sed -i '/shms.Shutdown()/d' main.go
sed -i '/shms := shm.NewServer(wss)/d' main.go
sed -i '/shms.StartToWatch()/d' main.go
sed -i 's/handleSignals(s, mcs, gs, wss, shms)/handleSignals(s, mcs, gs, wss)/g' main.go
sed -i 's/, shms \*shm.ShmServer//g' main.go

go build -o elibot-server ${API_V1}/*.go ${API_V2}/*.go ${API_V3}/*.go ${API}/*.go ${CONFIG}/*.go ${DB}/*.go ${LOG}/*.go ${MCSERVER}/*.go ${MIDDLEWARE}/*.go ${SERVERPB}/*.go ${SQLITEDB}/*.go ${VENDOR}/*.go