package userSys

import (
	"net/http"

	"github.com/eoscanada/eos-go"

	"github.com/gin-gonic/gin"
)

type Turn struct {
	Fuck string `json:"fuck"`
}

func (h *RestHandler) Register(c *gin.Context) {

	var userPk = "pk"
	var userAccount = "qwertasdffzxc"
	var to = "eosiotoken32"
	tx := &eos.PackedTransaction{}
	err := h.srvcContext.BCServer.EOSClint.PushTransaction()
	payload := &Turn{
		Fuck: "fuck asshole",
	}
	c.JSON(http.StatusOK, payload)
}
