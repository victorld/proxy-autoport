package api

import (
	"github.com/gin-gonic/gin"
	"proxy/server"
	"proxy/tools"
)

type MainApi struct {
}

func (mainApi *MainApi) ChangePort(c *gin.Context) {
	server.ChangePort()
	tools.Success(c, gin.H{"ret": "ok"}, "成功")

}
