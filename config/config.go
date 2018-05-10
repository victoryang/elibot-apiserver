package config

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

type GlobalConfiguration struct {
	AccessLogsFile		string		`yaml:"AccessLogsFile,omitempty"`
	ElibotLogsFile		string		`yaml:"ElibotLogsFile,omitempty"`
	EntryPoints			[]string	`yaml:"EntryPoints,omitempty"`
	ListenAddress		string		`yaml:"ListenAddress,omitempty"`
}

type ElibotServerConfiguration struct {
	GlobalConfig		*GlobalConfiguration
	ConfigFile			string
}

// Load parses the YAML input s into a Config.
func Load(s string) (*GlobalConfiguration, error) {
	cfg := &GlobalConfiguration{}

	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *ElibotServerConfiguration)LoadFile() error {
	content, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		return err
	}
	c.GlobalConfig, err = Load(string(content))
	if err != nil {
		return fmt.Errorf("parsing YAML file %s: %v", c.ConfigFile, err)
	}
	return nil
}

func NewConfiguration() *ElibotServerConfiguration {
	return &ElibotServerConfiguration {
			GlobalConfig:	&GlobalConfiguration{
				AccessLogsFile:			"/var/lib/elibot-server/access.Log",
				ElibotLogsFile:			"/var/lib/elibot-server/elibot.Log",
				EntryPoints:			[]string{"http"},
				ListenAddress:			"127.0.0.1:9000",
			},
			ConfigFile:	"/etc/elibot-server.yaml",
	}
}