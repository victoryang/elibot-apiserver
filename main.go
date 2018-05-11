package main

import (
	"os"
	"os/signal"
	"syscall"

	Log "elibot-apiserver/log"
	"elibot-apiserver/api"
	"elibot-apiserver/config"
)

func handleSignals(server *api.Server) {
	signal.Ignore()
	signalQueue := make(chan os.Signal)
	signal.Notify(signalQueue, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGUSR1)
	for {
		s := <-signalQueue
		switch s {
		case syscall.SIGHUP:
			// reload config file
			//helper.SetupConfig()
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
	c := config.NewConfiguration()
	/*err:=c.LoadFile() 
	if err!=nil {
		return
	}*/
	err := Log.OpenFile(c.GlobalConfig.ElibotLogsFile)
	if err!=nil {
		Log.Error("Could not open log file: ", err)
		return
	}
	defer Log.CloseFile()

	startAdminServer()
	s := api.NewApiServer(c.GlobalConfig)
	s.Run();
	handleSignals(s)
}