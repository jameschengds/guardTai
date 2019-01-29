package userSys

import (
	"private/context"

	logging "github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("handler")
)

// RestHandler escrow service rest handler
type RestHandler struct {
	srvcContext *context.Context
}

// NewRestHandler ...
func NewRestHandler(c *context.Context) (*RestHandler, error) {
	return &RestHandler{srvcContext: c}, nil
}
