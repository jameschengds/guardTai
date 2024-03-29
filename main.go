package main

import (
	set "github.com/jameschengds/guardTai/pkg/setting"

	"github.com/jameschengds/guardTai/server"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

func main() {

	cfg, err := set.GetServiceCfg() //读配置
	if err != nil {
		logger.Error("can't read config")
		return
	}
	ser, err := server.NewServer(cfg)
	if err != nil {
		logger.Panicf("Failed to create %s server, %+v", err)
		return
	}
	logger.Infof("Beginning to serve requests")
	//初始化
	ser.Start()
	//启动服务

}
