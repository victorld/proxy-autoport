package route

import (
	"github.com/gin-gonic/gin"
	"proxy/api"
	"proxy/tools"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	mainGroup := r.Group("/do", gin.BasicAuth(gin.Accounts{
		tools.HttpUser: tools.HttpPassword,
	}))

	var mainApi = new(api.MainApi)
	mainGroup.GET("change", mainApi.ChangePort)

	return r
}
