package server

import (
	"fmt"
	"private/context"
	"private/controler/userSys"
	set "private/pkg/setting"
	"private/routers"

	"github.com/op/go-logging"
)

var (
	Titan  *ServerImpl
	logger = logging.MustGetLogger("server")
)

type ServerImpl struct {
	config         *set.TitanSrvcConfig
	context        *context.Context
	userSysHandler *userSys.RestHandler
}

func (s *ServerImpl) Start() (err error) {
	// start to serve http connections
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
	return nil
}
func (s *ServerImpl) init() (err error) {
	s.context = context.GetServerContext()
	if nil == s.context {
		logger.Errorf("Server context is nil")
		return fmt.Errorf("context is nil")
	}
	s.context.Config = s.config
	err = s.context.Init()
	if nil != err {
		logger.Errorf("Initalize server context error: %v", err)
		return err
	}

	if err = s.httpHandlerInit(); err != nil {
		return err
	}
	return nil
}

func (s *ServerImpl) httpHandlerInit() (err error) {

	s.userSysHandler, err = userSys.NewRestHandler(s.context)
	if err != nil {
		logger.Errorf("Failed to create account rest http handler instance, %+v", err)
		return err
	}
	return nil
}
