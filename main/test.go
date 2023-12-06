package main

import (
	"proxy/cons"
	"proxy/plugin/clash"
	"proxy/tools"
)

func main() {
	tools.InitViper()
	cons.InitConst()
	tools.InitLogger()

	//server.ChangePort()
	clash.TestClash()
	//client.TestJob()

	//fmt.Println(url.QueryEscape("http://baidu.com/"))
}
