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
		EntryPoints:			[]string{"http"},
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
		return nil, err
	}
	return cfg, nil
}

func LoadFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg *Config
	cfg, err = Load(string(content))
	if err != nil {
		Log.Error("parsing YAML file %s: %s", filename, err)
		return nil, err
	}
	return cfg, nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// We want to set c to the defaults and then overwrite it with the input.
	// To make unmarshal fill the plain data struct rather than calling UnmarshalYAML
	// again, we have to hide it using a type indirection.
	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	if c.AccessLogsFile == nil {
		c.AccessLogsFile = DefaultConfig.AccessLogsFile
	}

	if c.ElibotLogsFile == nil {
		c.ElibotLogsFile = DefaultConfig.ElibotLogsFile
	}

	if c.EntryPoints == nil {
		c.EntryPoints = DefaultConfig.EntryPoints
	}

	if c.ListenAddress == nil {
		c.ListenAddress = DefaultConfig.ListenAddress
	}
	return nil
}