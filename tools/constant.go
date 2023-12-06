package tools

import (
	"fmt"
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
)

func InitConst() {
	HttpUser = GetConfigString("server.HttpUser")
	HttpPassword = GetConfigString("server.HttpPassword")
	ServerModifyLineNumber, _ = strconv.Atoi(GetConfigString("server.ServerModifyLineNumber"))
	DefaultPort, _ = strconv.Atoi(GetConfigString("server.DefaultPort"))
	ConfigFilePath = GetConfigString("server.ConfigFilePath")
	ChangePortShellPath = GetConfigString("server.ChangePortShellPath")
	ClashServer = GetConfigString("server.ClashServer")
	ClashSecret = GetConfigString("server.ClashSecret")
	GetPortShellPath = GetConfigString("server.GetPortShellPath")

	fmt.Println("HttpUser :", HttpUser)
	fmt.Println("HttpPassword :", HttpPassword)
	fmt.Println("ServerModifyLineNumber :", ServerModifyLineNumber)
	fmt.Println("DefaultPort :", DefaultPort)
	fmt.Println("ConfigFilePath :", ConfigFilePath)
	fmt.Println("ChangePortShellPath :", ChangePortShellPath)
	fmt.Println("ClashServer :", ClashServer)
	fmt.Println("ClashSecret :", ClashSecret)
	fmt.Println("GetPortShellPath :", GetPortShellPath)

}
