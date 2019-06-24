package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ServerImpl) SetupRouter() *gin.Engine {
	// Disable Console Color
	//// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLGlob("template/**")
	r.GET("/test", s.userSysHandler.Register)
	r.GET("/eos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template/eos.html", gin.H{
			"title": "Posts",
		})
	})

	return r
}
