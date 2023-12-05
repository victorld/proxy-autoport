package tools

import (
	"fmt"
	"strconv"
)

var (
	LineNumber     int
	DefaultPort    int
	HttpPassword   string
	ConfigFilePath string
	ShellPath      string
	HttpUser       string
)

func InitConst() {
	HttpUser = GetConfigString("server.HttpUser")
	HttpPassword = GetConfigString("server.HttpPassword")
	LineNumber, _ = strconv.Atoi(GetConfigString("server.LineNumber"))
	DefaultPort, _ = strconv.Atoi(GetConfigString("server.DefaultPort"))
	ConfigFilePath = GetConfigString("server.ConfigFilePath")
	ShellPath = GetConfigString("server.ShellPath")
	fmt.Println("HttpUser :", HttpUser)
	fmt.Println("HttpPassword :", HttpPassword)
	fmt.Println("LineNumber :", LineNumber)
	fmt.Println("DefaultPort :", DefaultPort)
	fmt.Println("ConfigFilePath :", ConfigFilePath)
	fmt.Println("ShellPath :", ShellPath)

}
