package cons

import (
	"fmt"
	"os"
	"proxy/tools"
	"strconv"
)

var (
	ServerModifyLineNumber    int
	ProxyDefaultPort          int
	HttpPassword              string
	ServerConfigFilePath      string
	ServerChangePortShellPath string
	HttpUser                  string
	ClashServer               string
	ClashSecret               string
	GetPortShellPath          string
	ServerHttpPort            string
	ClientLocalConfitFilePath string
	ClashServerConfitFilePath string
	ServerIP                  string
	TestUrl                   string
	ClientChangePortShellPath string
)

func InitConst() {
	//server
	HttpUser = tools.GetConfigString("server.HttpUser")
	HttpPassword = tools.GetConfigString("server.HttpPassword")
	ServerModifyLineNumber, _ = strconv.Atoi(tools.GetConfigString("server.ServerModifyLineNumber"))
	ProxyDefaultPort, _ = strconv.Atoi(tools.GetConfigString("server.ProxyDefaultPort"))
	ServerConfigFilePath = tools.GetConfigString("server.ServerConfigFilePath")
	ServerChangePortShellPath = tools.GetConfigString("server.ServerChangePortShellPath")
	GetPortShellPath = tools.GetConfigString("server.GetPortShellPath")
	ServerHttpPort = tools.GetConfigString("server.ServerHttpPort")
	//client
	ClashServer = tools.GetConfigString("client.ClashServer")
	ClashSecret = tools.GetConfigString("client.ClashSecret")
	TestUrl = tools.GetConfigString("client.TestUrl")
	ClientLocalConfitFilePath = tools.GetConfigString("client.ClientLocalConfitFilePath")
	ClashServerConfitFilePath = tools.GetConfigString("client.ClashServerConfitFilePath")
	ServerIP = tools.GetConfigString("client.ServerIP")
	ClientChangePortShellPath = tools.GetConfigString("client.ClientChangePortShellPath")

	if v := os.Getenv("client.ClashServer"); v != "" {
		ClashServer = v
	}
	if v := os.Getenv("client.ClashSecret"); v != "" {
		ClashSecret = v
	}
	if v := os.Getenv("client.ClientLocalConfitFilePath"); v != "" {
		ClientLocalConfitFilePath = v
	}
	if v := os.Getenv("client.ClashServerConfitFilePath"); v != "" {
		ClashServerConfitFilePath = v
	}
	if v := os.Getenv("client.ServerIP"); v != "" {
		ServerIP = v
	}
	if v := os.Getenv("client.ClientChangePortShellPath"); v != "" {
		ClientChangePortShellPath = v
	}
	fmt.Println("HttpUser :", HttpUser)
	fmt.Println("HttpPassword :", HttpPassword)
	fmt.Println("ServerModifyLineNumber :", ServerModifyLineNumber)
	fmt.Println("ProxyDefaultPort :", ProxyDefaultPort)
	fmt.Println("ServerConfigFilePath :", ServerConfigFilePath)
	fmt.Println("ServerChangePortShellPath :", ServerChangePortShellPath)
	fmt.Println("ClashServer :", ClashServer)
	fmt.Println("ClashSecret :", ClashSecret)
	fmt.Println("GetPortShellPath :", GetPortShellPath)
	fmt.Println("ServerHttpPort :", ServerHttpPort)
	fmt.Println("TestUrl :", TestUrl)
	fmt.Println("ClientLocalConfitFilePath :", ClientLocalConfitFilePath)
	fmt.Println("ClashServerConfitFilePath :", ClashServerConfitFilePath)
	fmt.Println("ServerIP :", ServerIP)
	fmt.Println("ClientChangePortShellPath :", ClientChangePortShellPath)
	fmt.Println()

}
