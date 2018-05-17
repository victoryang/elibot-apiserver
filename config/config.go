package config

import (
	"io/ioutil"

	Log "elibot-apiserver/log"
	"gopkg.in/yaml.v2"
)

var (
	DefaultDebug = &DEBUG {
		Enable:			false,
		Level:			4,
	}

	DefaultSqliteDB = &SqliteDB {
		Path:			"/root/",
		FileName:		"elibotDB",
	}

	DefaultBackup = &BackUp {
		Dir:			"backups/",
	}

	DefaultAdmin = &AdminSever {
		ListenAddress:			"0.0.0.0:9090",
	}

	DefaultConfig = &Config{
		AccessLogsFile:			"/var/lib/elibot-server/access.Log",
		ElibotLogsFile:			"/var/lib/elibot-server/elibot.Log",
		EntryPoints:			[]string{"http"},
		ListenAddress:			"0.0.0.0:9000",
		Sqlite:					DefaultSqliteDB,
		Backup:					DefaultBackup,
		Admin:					DefaultAdmin,
	}
)

type BackUp struct {
	Dir 				string		`yaml:"dir,omitempty"`
}

func (b *BackUp) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain BackUp
	if err := unmarshal((*plain)(b)); err != nil {
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
	type plain SqliteDB
	if err := unmarshal((*plain)(s)); err != nil {
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

type DEBUG struct {
	Enable			bool			`yaml:"enable,omitempty"`
	Level			int 			`yaml:"level,omitempty"`
}

func (d *DEBUG) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain DEBUG
	if err := unmarshal((*plain)(d)); err != nil {
		return err
	}

	if d.Enable == true {
		d.Level = 5
	} else {
		d.Level = DefaultDebug.Level
	}
	Log.SetLevelInt(d.Level)
	return nil
}

type AdminSever struct {
	ListenAddress		string		`yaml:"admin_address,omitempty"`
}

func (a *AdminSever) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain AdminSever
	if err := unmarshal((*plain)(a)); err != nil {
		return err
	}

	if a.ListenAddress == "" {
		a.ListenAddress = DefaultAdmin.ListenAddress
	}
	return nil
}


type Config struct {
	AccessLogsFile		string		`yaml:"accessLog,omitempty"`
	Debug				DEBUG		`yaml:"debug,omitempty"`
	ElibotLogsFile		string		`yaml:"serverLog,omitempty"`
	EntryPoints			[]string	`yaml:"entrypoints,omitempty"`
	ListenAddress		string		`yaml:"server_address,omitempty"`

	Sqlite				*SqliteDB	`yaml:"sqlite,omitempty"`
	Backup				*BackUp 	`yaml:"backup,omitempty"`
	Admin 				*AdminSever `yaml:"admin,omitempty"`
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

	if c.Admin == nil {
		c.Admin = DefaultAdmin
	}
	return nil
}