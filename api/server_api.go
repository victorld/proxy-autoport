package api

import (
	"github.com/gin-gonic/gin"
	"proxy/server"
	"proxy/tools"
)

type ServerApi struct {
}

func (serverApi *ServerApi) ChangePort(c *gin.Context) {
	port := server.ChangePort()
	tools.Success(c, gin.H{"port": port}, "成功")

}
