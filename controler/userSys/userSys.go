package userSys

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Turn struct {
	Fuck string `json:"fuck"`
}

func (h *RestHandler) Register(c *gin.Context) {
	payload := &Turn{
		Fuck: "fuck asshole",
	}
	c.JSON(http.StatusOK, payload)
}
