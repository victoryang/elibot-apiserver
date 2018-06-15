package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	Log "elibot-apiserver/log"
)

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

func (h *HttpEntryPoint) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain HttpEntryPoint
	if err := unmarshal((*plain)(h)); err != nil {
		return err
	}

	if h.ListenAddress == "" {
		h.ListenAddress = DefaultHttpEntryPoint.ListenAddress
	}
	return nil
}

func (e *GrpcEntryPoint) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain GrpcEntryPoint
	if err := unmarshal((*plain)(e)); err != nil {
		return err
	}

	if e.ListenAddress == "" {
		e.ListenAddress = DefaultGrpcEntryPoint.ListenAddress
	}
	return nil
}

func (w *WebsocketEntryPoint) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain WebsocketEntryPoint
	if err := unmarshal((*plain)(w)); err != nil {
		return err
	}

	if w.ListenAddress == "" {
		w.ListenAddress = DefaultWebsocketEntryPoint.ListenAddress
	}
	return nil
}

func (d *Database) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Database
	if err := unmarshal((*plain)(d)); err != nil {
		return err
	}
	if d.FileName == "" {
		s.FileName = DefaultDatabase.FileName
	}

	if d.BackupDir == "" {
		s.BackupDir = DefaultDatabase.BackupDir
	}
	return nil
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

func (g *GlobalConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain GlobalConfiguration
	if err := unmarshal((*plain)(g)); err != nil {
		return err
	}

	if g.AccessLogsFile == "" {
		g.AccessLogsFile = DefaultGlobalConfiguration.AccessLogsFile
	}

	if g.ServerLogsFile == "" {
		g.ServerLogsFile = DefaultGlobalConfiguration.ServerLogsFile
	}

	if g.Debug == nil {
		g.Debug = DefaultGlobalConfiguration.Debug
	}

	if g.HttpEntryPoint == nil {
		g.HttpEntryPoint = DefaultGlobalConfiguration.HttpEntryPoint
	}

	if g.GrpcEntryPoint == nil {
		g.GrpcEntryPoint = DefaultGlobalConfiguration.GrpcEntryPoint
	}

	if g.WebsocketEntryPoint == nil {
		g.WebsocketEntryPoint = DefaultGlobalConfiguration.WebsocketEntryPoint
	}

	if g.Databases == nil {
		g.Databases = DefaultGlobalConfiguration.Databases
	}

	if g.Admin == nil {
		g.Admin = DefaultGlobalConfiguration.Admin
	}
	return nil
}

// Load parses the YAML input s into a GlobalConfiguration.
func Load(s string) (*GlobalConfiguration, error) {
	cfg := &GlobalConfiguration{}

	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadFile(filename string) (*GlobalConfiguration, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg *GlobalConfiguration
	cfg, err = Load(string(content))
	if err != nil {
		return nil, err
	}
	return cfg, nil
}