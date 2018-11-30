package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"path"

	Log "elibot-apiserver/log"
	"elibot-apiserver/api"
	"elibot-apiserver/mcserver"
	"elibot-apiserver/paramserver"
	"elibot-apiserver/auth"
	"elibot-apiserver/sqlitedb"
	"elibot-apiserver/db"
	"elibot-apiserver/settings"
	"elibot-apiserver/config"
	"elibot-apiserver/shm"
	"elibot-apiserver/websocket"
	"elibot-apiserver/alarm"
	"github.com/spf13/cobra"
)

const (
	mcserverAddress = "127.0.0.1:8055"
)

func handleSignals(s *api.Server, gs *api.GrpcServer, wss *websocket.WsServer, shms *shm.ShmServer) error {
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

			db.CloseDB()

			Log.CloseFile()

			return nil
		}
	}
}

func ConfigServerLog(cfg *config.GlobalConfiguration) error {
	dir := path.Dir(cfg.ServerLogsFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
    	os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	}
	if err := Log.OpenFile(cfg.ServerLogsFile); err!=nil {
		Log.Error("Configure server log error: ", err)
		return err
	}

	Log.SetOwnFormatter("text")
	return nil
}

func SetUpDatabase(cfg *config.GlobalConfiguration) {
	dbcfg := cfg.Databases
	sqlitedb.SetupDB(dbcfg.FileName, dbcfg.BackupDir)
	sqlitedb.InitSqlitedb()
	db.SetupDB(dbcfg.FileName)
	settings.SetupTable()
}

func runDaemon(cmd *cobra.Command, args []string) error {
	cfg := LoadConfig()
	
	if err := ConfigServerLog(cfg); err!=nil {
		return returnError(ERR_OPEN_LOG_FILE_FAIL)
	}

	if err := auth.AuthInit(cfg.Secure); err!= nil {
		os.Exit(ERR_AUTH_INIT_FAIL)
	}

	startAdminServer(cfg.Admin)

	SetUpDatabase(cfg)

	if err := mcserver.NewRpcClient(); err!=nil {
		Log.Error("Error to connect to mc server")
		os.Exit(ERR_START_MCSERVER)
	}

	if err := paramserver.NewRpcClient(); err!=nil {
		Log.Error("Error to connect to param server")
		os.Exit(ERR_START_PARAMSERVER)
	}

	apiserver := api.NewApiServer(cfg)
	if apiserver == nil {
		Log.Error("Error in starting apiserver")
		return returnError(ERR_START_APISERVER)
	}
	apiserver.Run()
	
	gs := api.NewGrpcServer(cfg.Grpc)
	if gs == nil {
		Log.Error("Failed to start grpc server")
		return returnError(ERR_START_GRPCSERVER)
	}

	wss := websocket.NewWsServer(cfg.Websocket)
	wss.Run()

	if err := alarm.NewAlarmMonitor(wss); err!=nil {
		Log.Error("Could not watch server log")
	}

	shms,err := shm.NewServer(wss)
	if err!=nil {
		Log.Error(err.Error())
		return returnError(ERR_START_SHMSERVER)
	}
	shms.StartToWatch()
	return handleSignals(apiserver, gs, wss, shms)
}
