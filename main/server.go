package main

import (
	"github.com/gin-gonic/gin"
	"proxy/route"
	"proxy/tools"
)

func main() {

	tools.InitViper()
	tools.InitConst()
	r := gin.Default()
	r = route.InitRouter(r)
	r.Run(":37001")
}
