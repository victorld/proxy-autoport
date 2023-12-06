package main

import (
	"github.com/gin-gonic/gin"
	"proxy/cons"
	"proxy/route"
	"proxy/tools"
)

func main() {

	tools.InitViper()
	cons.InitConst()
	tools.InitLogger()

	r := gin.Default()
	r = route.InitRouter(r)
	r.Run(":" + cons.HttpPort)
}
