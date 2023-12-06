package cons

import (
	"fmt"
	"proxy/tools"
	"strconv"
)

var (
	ServerModifyLineNumber int
	ProxyDefaultPort       int
	HttpPassword           string
	ServerConfigFilePath   string
	ChangePortShellPath    string
	HttpUser               string
	ClashServer            string
	ClashSecret            string
	GetPortShellPath       string
	ServerHttpPort         string
	TestUrl                string
)

func InitConst() {
	//server
	HttpUser = tools.GetConfigString("server.HttpUser")
	HttpPassword = tools.GetConfigString("server.HttpPassword")
	ServerModifyLineNumber, _ = strconv.Atoi(tools.GetConfigString("server.ServerModifyLineNumber"))
	ProxyDefaultPort, _ = strconv.Atoi(tools.GetConfigString("server.ProxyDefaultPort"))
	ServerConfigFilePath = tools.GetConfigString("server.ServerConfigFilePath")
	ChangePortShellPath = tools.GetConfigString("server.ChangePortShellPath")
	GetPortShellPath = tools.GetConfigString("server.GetPortShellPath")
	ServerHttpPort = tools.GetConfigString("server.ServerHttpPort")
	//client
	ClashServer = tools.GetConfigString("client.ClashServer")
	ClashSecret = tools.GetConfigString("client.ClashSecret")
	TestUrl = tools.GetConfigString("client.TestUrl")

	fmt.Println("HttpUser :", HttpUser)
	fmt.Println("HttpPassword :", HttpPassword)
	fmt.Println("ServerModifyLineNumber :", ServerModifyLineNumber)
	fmt.Println("ProxyDefaultPort :", ProxyDefaultPort)
	fmt.Println("ServerConfigFilePath :", ServerConfigFilePath)
	fmt.Println("ChangePortShellPath :", ChangePortShellPath)
	fmt.Println("ClashServer :", ClashServer)
	fmt.Println("ClashSecret :", ClashSecret)
	fmt.Println("GetPortShellPath :", GetPortShellPath)
	fmt.Println("ServerHttpPort :", ServerHttpPort)
	fmt.Println("TestUrl :", TestUrl)
	fmt.Println()

}
