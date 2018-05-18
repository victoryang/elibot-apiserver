package main

import (
	"os"
	"os/signal"
	"syscall"

	Log "elibot-apiserver/log"
	"elibot-apiserver/api"
	"elibot-apiserver/config"
)

const (
	configFile = "/rbctrl/configuration/elibot-server.yaml"
)

func handleSignals(server *api.Server) {
	signal.Ignore()
	signalQueue := make(chan os.Signal)
	signal.Notify(signalQueue, syscall.SIGHUP, os.Interrupt)
	for {
		s := <-signalQueue
		switch s {
		//TODO:
		//case syscall.SIGHUP:
			//reload config file
		default:
			// stop server
			stopAdminServer()
			server.Shutdown()
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

	startAdminServer(cfg)
	s := api.NewApiServer(cfg)
	s.Run();
	handleSignals(s)
}