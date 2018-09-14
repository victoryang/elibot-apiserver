package cmd

import(
	"elibot-apiserver/config"
	Log "elibot-apiserver/log"
)

var configFile string

func LoadConfig() *config.GlobalConfiguration{
	cfg, err := config.LoadFile(configFile) 
	if err!=nil {
		Log.Error("Parse configure file error: ", err)
		Log.Error("set to default configuration")
		cfg = config.DefaultGlobalConfiguration
	}

	return cfg
}