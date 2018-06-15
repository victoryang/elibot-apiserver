package main

import (
	"os"
	"os/signal"
	"syscall"

	Log "elibot-apiserver/log"
	"elibot-apiserver/api"
	"elibot-apiserver/dbproxy"
	"elibot-apiserver/config"
	"elibot-apiserver/mcserver"
	"elibot-apiserver/shm"
)

const (
	configFile = "/rbctrl/configuration/elibot-server.yaml"
	mcserverAddress = "127.0.0.1:8055"
)

const (
	ERR_NONE = iota
	ERR_OPEN_LOG_FILE_FAIL
	ERR_START_MCSERVER
	ERR_START_APISERVER
	ERR_START_GRPCSERVER
	ERR_START_WSSERVER
	ERR_START_SHMSERVER
)

func handleSignals(s *api.Server, mcs *mcserver.MCserver, gs *api.GrpcServer, wss *api.WsServer, shms *shm.ShmServer) {
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

func LoadConfig() *config.GlobalConfiguration{
	cfg, err := config.LoadFile(configFile) 
	if err!=nil {
		Log.Error("Parse configure file error: ", err)
		Log.Error("set to default configuration")
		cfg = config.DefaultGlobalConfiguration
	}

	return cfg
}

func ConfigServerLog(cfg *config.GlobalConfiguration) error {
	if err := Log.OpenFile(cfg.ServerLogsFile); err!=nil {
		Log.Error("Configure server log error: ", err)
		return err
	}

	Log.SetOwnFormatter("text")
	return nil
}

func SetUpDatabase(cfg *config.GlobalConfiguration) {
	db := cfg.Databases
	dbproxy.DBsetup(db.FileName, db.BackupDir)
}

func main() {
	cfg := LoadConfig()
	
	if err := ConfigServerLog(cfg); err!=nil {
		os.Exit(ERR_OPEN_LOG_FILE_FAIL)
	}
	defer Log.CloseFile()

	mcs := mcserver.NewMCServer(mcserverAddress, 3)
	if mcs==nil {
		Log.Error("Error in connecting to mc server")
		os.Exit(ERR_START_MCSERVER)
	}

	startAdminServer(cfg.Admin)

	SetUpDatabase(cfg)
	s := api.NewApiServer(cfg)
	s.Run()
	
	gs := api.NewGrpcServer(cfg.Grpc)
	if gs == nil {
		Log.Error("Failed to start grpc server")
		os.Exit(ERR_START_GRPCSERVER)
	}

	wss := api.NewWsServer(cfg.Websocket)
	wss.Run()

	shms,err := shm.NewServer(wss)
	if err!=nil {
		Log.Error(err.Error())
		os.Exit(ERR_START_SHMSERVER)
	}
	shms.StartToWatch()
	handleSignals(s, mcs, gs, wss, shms)
}