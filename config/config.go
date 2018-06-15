package config

import (
)

var (
	DefaultDebug = &DEBUG {
		Enable:			false,
		Level:			4,
	}

	DefaultHttpEntryPoint = &HttpEntryPoint {
		ListenAddress:			"0.0.0.0:9000",
	}

	DefaultGrpcEntryPoint = &GrpcEntryPoint {
		ListenAddress:			"0.0.0.0:9500",
	}

	DefaultWebsocketEntryPoint = &WebsocketEntryPoint {
		ListenAddress:			"0.0.0.0:9050",
	}

	DefaultDatabase = &Database {
		FileName:		"/rbctrl/db/elibotDB.db",
		BackupDir:		"/rbctrl/db/backups/",
	}

	DefaultAdmin = &AdminSever {
		ListenAddress:			"0.0.0.0:9090",
	}

	DefaultGlobalConfiguration = &GlobalConfiguration{
		AccessLogsFile:			"/rbctrl/apiserver/log/access.Log",
		ServerLogsFile:			"/rbctrl/apiserver/log/server.Log",
		Debug:					DefaultDebug,

		HttpEntryPoint:			DefaultHttpEntryPoint,
		GrpcEntryPoint: 		DefaultGrpcEntryPoint,
		WebsocketEntryPoint:	DefaultWebsocketEntryPoint,

		Databases:		DefaultDatabase,
		Admin:			DefaultAdmin,
	}
)

type DEBUG struct {
	Enable			bool			`yaml:"enable,omitempty"`
	Level			int 			`yaml:"level,omitempty"`
}

type HttpEntryPoint struct {
	ListenAddress	string			`yaml:"address,omitempty"`
}

type GrpcEntryPoint struct {
	ListenAddress	string			`yaml:"address,omitempty"`
}

type WebsocketEntryPoint struct {
	ListenAddress	string			`yaml:"address,omitempty"`
}

type Database struct {
	FileName		string			`yaml:"filename,omitempty"`
	BackupDir		string			`yaml:"backupdir,omitempty"`
}

type AdminSever struct {
	ListenAddress		string		`yaml:"admin_address,omitempty"`
}

type GlobalConfiguration struct {
	AccessLogsFile		string					`yaml:"accesslog,omitempty"`
	ServerLogsFile		string					`yaml:"serverlog,omitempty"`
	Debug				*DEBUG					`yaml:"debug,omitempty"`

	Http 				*HttpEntryPoint 				`yaml:"http,omitempty`
	Grpc 				*GrpcEntryPoint 				`yaml:"grpc,omitempty`
	Websocket 			*WebsocketEntryPoint 		`yaml:"websocket,omitempty`

	Databases			*Database				`yaml:"databases,omitempty"`
	Admin 				*AdminSever 			`yaml:"admin,omitempty"`
}