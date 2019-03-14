package server

import "guardTai/pkg/setting"

type Server interface {
	Start() (err error)
	//Shutdown()
}

// NewServer create escrow-service server
func NewServer(c *setting.TitanSrvcConfig) (Server, error) {
	if Titan == nil {
		Titan = &ServerImpl{
			config: c,
		}
		err := Titan.init()
		if err != nil {
			logger.Errorf("Failed to initialize escrow server, %+v", err)
			return nil, err
		}
	}
	return Titan, nil
}
