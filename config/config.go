package config

import (
	"io/ioutil"

	Log "elibot-apiserver/log"
	"gopkg.in/yaml.v2"
)

var (
	DefaultSqliteDB = &SqliteDB {
		Path:			"/root/",
		FileName:		"elibotDB",
	}

	DefaultBackup = &BackUp {
		Dir:			"backups/",
	}

	DefaultConfig = &Config{
		AccessLogsFile:			"/var/lib/elibot-server/access.Log",
		ElibotLogsFile:			"/var/lib/elibot-server/elibot.Log",
		EntryPoints:			[]string{"http"},
		ListenAddress:			"127.0.0.1:9000",
		Sqlite:					DefaultSqliteDB,
		Backup:					DefaultBackup,
	}
)

type BackUp struct {
	Dir 				string		`yaml:"dir,omitempty"`
}

func (b *Backup) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	if b.Dir == "" {
		b.Dir = DefaultBackup.Dir
	}
	return nil
}

type SqliteDB struct {
	Path 				string		`yaml:"path,omitempty"`
	FileName 			string		`yaml:"name,omitempty"`
}

func (s *SqliteDB) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}
	if s.Path == "" {
		s.Path = DefaultSqliteDB.Path
	}

	if s.FileName == "" {
		s.FileName = DefaultSqliteDB.FileName
	}
	return nil
}

type Config struct {
	AccessLogsFile		string		`yaml:"accessLog,omitempty"`
	ElibotLogsFile		string		`yaml:"serverLog,omitempty"`
	EntryPoints			[]string	`yaml:"entrypoints,omitempty"`
	ListenAddress		string		`yaml:"server_address,omitempty"`

	Sqlite				*SqliteDB	`yaml:"sqlite,omitempty"`
	Backup				*BackUp 	`yaml:"backup,omitempty"`
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

	if c.AccessLogsFile == "" {
		c.AccessLogsFile = DefaultConfig.AccessLogsFile
	}

	if c.ElibotLogsFile == "" {
		c.ElibotLogsFile = DefaultConfig.ElibotLogsFile
	}

	if c.EntryPoints == nil {
		c.EntryPoints = DefaultConfig.EntryPoints
	}

	if c.ListenAddress == "" {
		c.ListenAddress = DefaultConfig.ListenAddress
	}

	if c.Sqlite == nil {
		c.Sqlite = DefaultSqliteDB
	}

	if c.Backup == nil {
		c.Backup = DefaultBackup
	}
	return nil
}