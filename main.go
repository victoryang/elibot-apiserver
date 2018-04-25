package main

import (
	"os"
	"os/signal"
	"syscall"
	"api"
)

func func main() {
	//eliBotConfig := config.NewElibotConfig()
	//admin := api.NewAdminServer()
	config := new(api.ServerConfig)
	config.addr = "www.elibot.cn:9000"
	s := api.NewApiServer(config)
	s.Run();

	// ignore signal handlers set by Iris
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
			// stop YIG server, order matters
			s.Stop()
			os.Exit(0)
			return
		}
	}
}