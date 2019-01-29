package context

import (
	"private/models"
	"private/pkg/setting"

	"github.com/op/go-logging"
)

var (
	serverContext *Context
	logger        = logging.MustGetLogger("context")
)

type Context struct {
	Config  *setting.TitanSrvcConfig
	Storage models.DBBackend
}

func GetServerContext() *Context {
	if serverContext == nil {
		serverContext = &Context{}
	}
	return serverContext
}

// Init ...
func (c *Context) Init() error {
	if err := c.storageInit(); err != nil {
		logger.Errorf("Failed to initialize storage backend: %+v", err)
		return err
	}
	return nil
}

func (c *Context) storageInit() error {
	var err error
	c.Storage, err = models.NewDbBackend(&c.Config.Database)
	if err != nil {
		logger.Errorf("New database backend error, %+v", err)
		return err
	}
	return nil
}
