package api

import (
	"github.com/gin-gonic/gin"
	"proxy/server"
	"proxy/tools"
)

type ServerApi struct {
}

func (serverApi *ServerApi) IncrPort(c *gin.Context) {
	port := server.IncrPort()
	tools.Success(c, gin.H{"port": port}, "成功")

}

func (serverApi *ServerApi) ChangePort(c *gin.Context) {
	portStr, ok := c.GetQuery("port")
	if ok {
		portNew, err := server.ChangePort(portStr)
		if err != nil {
			tools.Fail(c, "失败，原因："+err.Error(), nil)
		} else {
			tools.Success(c, gin.H{"port": portNew}, "成功")
		}
	} else {
		tools.Fail(c, "失败，原因：port参数读取失败", nil)
	}

}

func (serverApi *ServerApi) GetPort(c *gin.Context) {
	port := server.GetPort()
	tools.Success(c, gin.H{"port": port}, "成功")

}
