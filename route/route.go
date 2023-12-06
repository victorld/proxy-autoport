package route

import (
	"github.com/gin-gonic/gin"
	"proxy/api"
	"proxy/tools"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	serverGroup := r.Group("/server", gin.BasicAuth(gin.Accounts{
		tools.HttpUser: tools.HttpPassword,
	}))

	var serverApi = new(api.ServerApi)
	serverGroup.GET("change-port", serverApi.ChangePort)
	serverGroup.GET("get-port", serverApi.GetPort)

	return r
}
