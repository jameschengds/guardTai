package server

import (
	"github.com/gin-gonic/gin"
)

func (s *ServerImpl) SetupRouter() *gin.Engine {
	// Disable Console Color
	//// gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/test", s.userSysHandler.Register)

	return r
}
