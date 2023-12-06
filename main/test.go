package main

import (
	"fmt"
	"proxy/cons"
	"proxy/plugin/clash"
	"proxy/tools"
)

func testClash() {
	proxies, err := clash.GetProxies()
	if err != nil {
	}
	fmt.Println("proxies : ", proxies)
	fmt.Println()

	config, err := clash.GetConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("config : ", config)
	fmt.Println()

	//rules, err := clash.GetRules()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("rules : ", rules)
	//fmt.Println()

	delay, err := clash.GetProxyDelay("xg", "http:%2F%2Fwww.gstatic.com%2Fgenerate_204", 10000)
	if err != nil {
		panic(err)
	}
	fmt.Println("delay : ", delay)
	fmt.Println()
}

func main() {
	tools.InitViper()
	cons.InitConst()
	//server.ChangePort()
	testClash()
}
