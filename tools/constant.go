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
	ShellPath              string
	HttpUser               string
	ClashServer            string
	ClashSecret            string
)

func InitConst() {
	HttpUser = GetConfigString("server.HttpUser")
	HttpPassword = GetConfigString("server.HttpPassword")
	ServerModifyLineNumber, _ = strconv.Atoi(GetConfigString("server.ServerModifyLineNumber"))
	DefaultPort, _ = strconv.Atoi(GetConfigString("server.DefaultPort"))
	ConfigFilePath = GetConfigString("server.ConfigFilePath")
	ShellPath = GetConfigString("server.ShellPath")
	ClashServer = GetConfigString("server.ClashServer")
	ClashSecret = GetConfigString("server.ClashSecret")
	fmt.Println("HttpUser :", HttpUser)
	fmt.Println("HttpPassword :", HttpPassword)
	fmt.Println("ServerModifyLineNumber :", ServerModifyLineNumber)
	fmt.Println("DefaultPort :", DefaultPort)
	fmt.Println("ConfigFilePath :", ConfigFilePath)
	fmt.Println("ShellPath :", ShellPath)
	fmt.Println("ClashServer :", ClashServer)
	fmt.Println("ClashSecret :", ClashSecret)

}
