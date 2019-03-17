package setting

import (
	"log"
	"time"

	logging "github.com/op/go-logging"
	"gopkg.in/ini.v1"
)

var (
	logger = logging.MustGetLogger("config")
)

type TitanSrvcConfig struct {
	Cfg          *ini.File
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	JwtSecret    string
	Database     Database
	EOS          Eos
}
type Database struct {
	Host         string
	Type         string
	Driver       string
	Address      string
	DBname       string
	User         string
	Password     string
	SSLMode      string
	Enabled      bool
	MaxIdleConns int
	MaxConns     int
}
type Eos struct {
	PRC_SERVE string
}

func GetServiceCfg() (*TitanSrvcConfig, error) {
	var err error
	getServiceCfg := &TitanSrvcConfig{}
	getServiceCfg.Cfg, err = ini.Load("conf/app.ini")
	err = getServiceCfg.LoadBaseCfg()
	if err != nil {
		logger.Errorf("Failed to LoadBaseCfg config: %v", err)
		return nil, err
	}
	err = getServiceCfg.LoadServer()
	if err != nil {
		logger.Errorf("Failed to LoadServer config: %v", err)
		return nil, err
	}
	err = getServiceCfg.LoadApp()
	if err != nil {
		logger.Errorf("Failed to LoadApp config: %v", err)
		return nil, err
	}
	err = getServiceCfg.LoadDB()
	if err != nil {
		logger.Errorf("Failed to LoadDB config: %v", err)
		return nil, err
	}
	err = getServiceCfg.LoadEOS()
	if err != nil {
		logger.Errorf("Failed toLoadEOS config: %v", err)
		return nil, err
	}
	return getServiceCfg, nil
}

func (c *TitanSrvcConfig) LoadBaseCfg() error {
	c.RunMode = c.Cfg.Section("").Key("RUN_MODE").MustString("debug")
	return nil
}

func (c *TitanSrvcConfig) LoadServer() error {
	sec, err := c.Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
		return err
	}

	c.HTTPPort = sec.Key("HTTP_PORT").MustString("8000")
	c.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	c.WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	return nil
}

func (c *TitanSrvcConfig) LoadApp() error {
	sec, err := c.Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
		return err
	}

	c.JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	return nil
}

func (c *TitanSrvcConfig) LoadDB() error {
	sec, err := c.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
		return err
	}
	c.Database.Address = sec.Key("HOST").String()
	c.Database.Type = sec.Key("TYPE").String()
	c.Database.DBname = sec.Key("NAME").String()
	c.Database.User = sec.Key("USER").String()
	c.Database.Host = sec.Key("HOST").String()
	c.Database.Password = sec.Key("PASSWORD").String()

	return nil
}

func (c *TitanSrvcConfig) LoadEOS() error {
	sec, err := c.Cfg.GetSection("eos")
	if err != nil {
		log.Fatalf("Fail to get section 'eos': %v", err)
		return err
	}
	c.EOS.PRC_SERVE = sec.Key("PRC_SERVE").String()

	return nil
}
