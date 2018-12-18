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
		if d.Level == 0 {
			d.Level = 5
		}	
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

	if w.IndexFile == "" {
		w.IndexFile = DefaultWebsocketEntryPoint.IndexFile
	}
	return nil
}

func (d *Database) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Database
	if err := unmarshal((*plain)(d)); err != nil {
		return err
	}
	if d.FileName == "" {
		d.FileName = DefaultDatabase.FileName
	}

	if d.BackupDir == "" {
		d.BackupDir = DefaultDatabase.BackupDir
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

func (j *JwtToken) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain JwtToken
	if err := unmarshal((*plain)(j)); err != nil {
		return err
	}

	if j.PrivateKey == "" {
		j.PrivateKey = DefaultJwtToken.PrivateKey
	}

	if j.Signature == "" {
		j.Signature = DefaultJwtToken.Signature
	}
	return nil
}

func (c *Certificate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Certificate
	if err := unmarshal((*plain)(c)); err != nil {
		return err
	}

	if c.Path == "" {
		c.Path = DefaultCertificate.Path
	}

	if c.CaCert == "" {
		c.CaCert = DefaultCertificate.CaCert
	}
	
	if c.Certfile == "" {
		c.Certfile = DefaultCertificate.Certfile
	}

	if c.Keyfile == "" {
		c.Keyfile = DefaultCertificate.Keyfile
	}
	return nil
}

func (s *Security) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Security
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	if s.SSLCert == nil {
		s.SSLCert = DefaultSecurity.SSLCert
	}

	if s.Jwt == nil {
		s.Jwt = DefaultSecurity.Jwt
	}

	s.WhiteList = append(s.WhiteList, DefaultSecurity.WhiteList...)
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

	if g.UploadPath == "" {
		g.UploadPath = DefaultGlobalConfiguration.UploadPath
	}

	if g.Debug == nil {
		g.Debug = DefaultGlobalConfiguration.Debug
	}

	if g.Http == nil {
		g.Http = DefaultGlobalConfiguration.Http
	}

	if g.Grpc == nil {
		g.Grpc = DefaultGlobalConfiguration.Grpc
	}

	if g.Websocket == nil {
		g.Websocket = DefaultGlobalConfiguration.Websocket
	}

	if g.Databases == nil {
		g.Databases = DefaultGlobalConfiguration.Databases
	}

	if g.Admin == nil {
		g.Admin = DefaultGlobalConfiguration.Admin
	}

	if g.Secure == nil {
		g.Secure = DefaultGlobalConfiguration.Secure
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