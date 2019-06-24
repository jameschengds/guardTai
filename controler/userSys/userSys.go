package userSys

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameschengds/guardTai/api"
)

type Turn struct {
	Fuck string `json:"fuck"`
}

func (h *RestHandler) Register(c *gin.Context) {
	req := &api.PushTX{}
	if err := c.BindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, "param error")
		return
	}
	id, err := h.srvcContext.BCServer.PushTX(req.From, req.To, req.Pk, req.Memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "sign error")
		return
	}
	c.JSON(http.StatusOK, id)
}

//func (h *RestHandler) NewAccount(c *gin.Context) {
//	req := &api.NewAccountReq{}
//
//	if err := c.BindJSON(req); err != nil {
//		c.JSON(http.StatusBadRequest, "param error")
//		return
//	}
//	//id, err := h.srvcContext.BCServer.NewAccount()
//
//}
