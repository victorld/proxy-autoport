package route

import (
	"github.com/gin-gonic/gin"
	"proxy/api"
	"proxy/cons"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	serverGroup := r.Group("/server", gin.BasicAuth(gin.Accounts{
		cons.HttpUser: cons.HttpPassword,
	}))

	var serverApi = new(api.ServerApi)
	serverGroup.GET("change-port", serverApi.ChangePort)
	serverGroup.GET("get-port", serverApi.GetPort)
	serverGroup.GET("incr-port", serverApi.IncrPort)

	return r
}
