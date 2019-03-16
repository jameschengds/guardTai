package server

import (
	"fmt"
	"guardTai/blockchain"
	"guardTai/context"
	"guardTai/controler/userSys"
	set "guardTai/pkg/setting"

	"github.com/op/go-logging"
)

var (
	Titan  *ServerImpl
	logger = logging.MustGetLogger("server")
)

type ServerImpl struct {
	config         *set.TitanSrvcConfig
	context        *context.Context
	eosClient      *blockchain.Eos
	userSysHandler *userSys.RestHandler
}

func (s *ServerImpl) Start() (err error) {
	// start to serve http connections
	r := s.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run()
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
	if err = s.blockchainInit(); err != nil {
		logger.Errorf("Initalize blockchainInit: %v", err)
		return err
	}
	if err = s.httpHandlerInit(); err != nil {
		logger.Errorf("Initalize httpHandlerIniterror: %v", err)
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

func (s *ServerImpl) blockchainInit() (err error) {
	s.eosClient = blockchain.EOSCilentInit(s.config.Eos.PRC_SERVE)
	logger.Debugf("init eos client success")
	return nil
}
