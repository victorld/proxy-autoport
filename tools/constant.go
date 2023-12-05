package tools

import "strconv"

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

}
