package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"path"

	Log "elibot-apiserver/log"
	"elibot-apiserver/api"
	"elibot-apiserver/mcserver"
	"elibot-apiserver/auth"
	"elibot-apiserver/sqlitedb"
	"elibot-apiserver/db"
	"elibot-apiserver/settings"
	"elibot-apiserver/config"
	"elibot-apiserver/resource"
	"elibot-apiserver/websocket"
	"elibot-apiserver/alarm"
	"elibot-apiserver/netlink"
	"github.com/spf13/cobra"
)

const (
	mcserverAddress = "127.0.0.1:8055"
)

func handleSignals(s *api.Server, gs *api.GrpcServer, nlsServer *netlink.NLServer, reserver *resource.ResServer) error {
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
			reserver.Shutdown()

			if nlsServer!=nil {
				nlsServer.Close()
			}

			websocket.Shutdown()

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
	sqlitedb.InitSqliteDB(dbcfg.FileName)
	db.SetupDB(dbcfg.FileName)
	settings.SetupTable()
}

func runDaemon(cmd *cobra.Command, args []string) error {
	cfg := LoadConfig()
	
	if err := ConfigServerLog(cfg); err!=nil {
		return returnError(ERR_OPEN_LOG_FILE_FAIL)
	}

	startAdminServer(cfg.Admin)

	SetUpDatabase(cfg)

	if err := auth.AuthInit(cfg.Secure); err!= nil {
		Log.Error("Error to init auth module")
		os.Exit(ERR_AUTH_INIT_FAIL)
	}

	if err := mcserver.NewRpcClient(); err!=nil {
		Log.Error("Error to connect to mc server")
		os.Exit(ERR_START_MCSERVER)
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

	websocket.StartServer(cfg.Websocket)

	if err := alarm.NewAlarmMonitor(); err!=nil {
		Log.Error("Could not watch server log")
	}

	nlsServer := netlink.NewNetlinkServer()
	if nlsServer == nil {
		Log.Error("Error in starting nlsserver")
	}
	nlsServer.Start()

	reserver,err := resource.NewServer()
    if err != nil {
        os.Exit(ERR_START_RESERVER)
    }
    reserver.Start()
	return handleSignals(apiserver, gs, nlsServer, reserver)
}
