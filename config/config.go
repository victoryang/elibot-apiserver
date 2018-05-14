package config

import (
	"io/ioutil"

	Log "elibot-apiserver/log"
	"gopkg.in/yaml.v2"
)

var (
	DefaultConfig = &Config{
		AccessLogsFile:			"/var/lib/elibot-server/access.Log",
		ElibotLogsFile:			"/var/lib/elibot-server/elibot.Log",
		EntryPoints:			nil,
		ListenAddress:			"127.0.0.1:9000",
	}
)

type Config struct {
	AccessLogsFile		string		`yaml:"AccessLogsFile,omitempty"`
	ElibotLogsFile		string		`yaml:"ElibotLogsFile,omitempty"`
	EntryPoints			[]string	`yaml:"EntryPoints,omitempty"`
	ListenAddress		string		`yaml:"ListenAddress,omitempty"`
}

// Load parses the YAML input s into a Config.
func Load(s string) (*Config, error) {
	cfg := &Config{}

	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		Log.Error("parsing YAML file %s: %s", filename, err)
		return nil, err
	}
	return cfg, nil
}

func LoadFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return Load(string(content))
}