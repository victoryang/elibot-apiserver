package main

import (
	"os"
	"os/signal"
	"syscall"

	Log "elibot-apiserver/log"
	"elibot-apiserver/api"
	"elibot-apiserver/config"
	"elibot-apiserver/mcserver"
	"elibot-apiserver/shm"
)

const (
	configFile = "/rbctrl/configuration/elibot-server.yaml"
)

func handleSignals(s *api.Server, mcs *mcserver.MCserver, gs *api.GrpcServer, wss *api.WsServer, shms shm.ShmServer) {
	signal.Ignore()
	signalQueue := make(chan os.Signal)
	signal.Notify(signalQueue, syscall.SIGHUP, os.Interrupt)
	for {
		sig := <-signalQueue
		switch sig {
		//TODO:
		//case syscall.SIGHUP:
			//reload config file
		default:
			// stop server
			shms.Shutdown()

			wss.Shutdown()

			gs.Shutdown()
			
			s.Shutdown()
			stopAdminServer()

			mcs.Close()

			os.Exit(0)
			return
		}
	}
}

func main() {
	cfg, err := config.LoadFile(configFile) 
	if err!=nil {
		Log.Error("Parse configure file error: ", err)
		cfg = config.DefaultConfig
	}
	Log.Print("EntryPoints: ", cfg.EntryPoints)
	err = Log.OpenFile(cfg.RootDir+cfg.ElibotLogsFile)
	if err!=nil {
		Log.Error("Error in opening file ", cfg.RootDir+cfg.ElibotLogsFile, " ", err)
		return
	}
	defer Log.CloseFile()
	Log.SetOwnFormatter("text")

	mcs := mcserver.NewMCServer("127.0.0.1:8055", 3)
	if mcs==nil {
		Log.Error("Error in connecting to mc server")
	}

	startAdminServer(cfg)
	s := api.NewApiServer(cfg)
	s.Run()
	
	gs := api.NewGrpcServer()

	wss := api.NewWsServer()
	wss.Run()

	shms := shm.NewAndWatch(wss)
	shms.StartToWatch()
	handleSignals(s, mcs, gs, wss, shms)
}