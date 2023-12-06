package cons

import (
	"fmt"
	"proxy/tools"
	"strconv"
)

var (
	ServerModifyLineNumber int
	DefaultPort            int
	HttpPassword           string
	ConfigFilePath         string
	ChangePortShellPath    string
	HttpUser               string
	ClashServer            string
	ClashSecret            string
	GetPortShellPath       string
	HttpPort               string
)

func InitConst() {
	//server
	HttpUser = tools.GetConfigString("server.HttpUser")
	HttpPassword = tools.GetConfigString("server.HttpPassword")
	ServerModifyLineNumber, _ = strconv.Atoi(tools.GetConfigString("server.ServerModifyLineNumber"))
	DefaultPort, _ = strconv.Atoi(tools.GetConfigString("server.DefaultPort"))
	ConfigFilePath = tools.GetConfigString("server.ConfigFilePath")
	ChangePortShellPath = tools.GetConfigString("server.ChangePortShellPath")
	GetPortShellPath = tools.GetConfigString("server.GetPortShellPath")
	HttpPort = tools.GetConfigString("server.HttpPort")
	//client
	ClashServer = tools.GetConfigString("client.ClashServer")
	ClashSecret = tools.GetConfigString("client.ClashSecret")

	fmt.Println("HttpUser :", HttpUser)
	fmt.Println("HttpPassword :", HttpPassword)
	fmt.Println("ServerModifyLineNumber :", ServerModifyLineNumber)
	fmt.Println("DefaultPort :", DefaultPort)
	fmt.Println("ConfigFilePath :", ConfigFilePath)
	fmt.Println("ChangePortShellPath :", ChangePortShellPath)
	fmt.Println("ClashServer :", ClashServer)
	fmt.Println("ClashSecret :", ClashSecret)
	fmt.Println("GetPortShellPath :", GetPortShellPath)
	fmt.Println("HttpPort :", HttpPort)

}
