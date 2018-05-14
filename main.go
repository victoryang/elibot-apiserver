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
	configFile = "/etc/elibot-server.yaml"
)

func handleSignals(server *api.Server) {
	signal.Ignore()
	signalQueue := make(chan os.Signal)
	signal.Notify(signalQueue, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGUSR1)
	for {
		s := <-signalQueue
		switch s {
		//TODO:
		//case syscall.SIGHUP:
			//reload config file
			
		case syscall.SIGUSR1:
			//go DumpStacks()
		default:
			// stop server
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
	err = Log.OpenFile(cfg.ElibotLogsFile)
	if err!=nil {
		Log.Error("Could not open log file: ", err)
		return
	}
	defer Log.CloseFile()

	startAdminServer()
	s := api.NewApiServer(cfg)
	s.Run();
	handleSignals(s)
}