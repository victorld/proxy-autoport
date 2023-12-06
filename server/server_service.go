package server

import (
	"os/exec"
	"proxy/cons"
	"proxy/tools"
	"strconv"
	"strings"
)

func ChangePort() string {
	content := tools.ReadFileLine(cons.ServerModifyLineNumber, cons.ConfigFilePath)
	split := strings.TrimSpace(strings.Split(strings.Split(content, ":")[1], ",")[0])
	port, _ := strconv.Atoi(split)
	if port == 0 || port >= cons.DefaultPort+3000 {
		port = cons.DefaultPort
	}
	newPort := port + 1
	tools.Logger.Info("newPort : ", newPort)
	command := exec.Command("/bin/sh", "-c", cons.ChangePortShellPath+" "+strconv.Itoa(newPort)) //初始化Cmd
	bytes, err := command.Output()
	if err != nil {
		tools.Logger.Error("shell error : ", err)
		return ""
	}

	tools.Logger.Info("shell out : ", string(bytes))
	return strconv.Itoa(newPort)
}

func GetPort() string {
	command := exec.Command("/bin/sh", "-c", cons.GetPortShellPath) //初始化Cmd
	bytes, err := command.Output()
	if err != nil {
		tools.Logger.Error("shell error : ", err)
		return ""
	}

	tools.Logger.Info("shell out : ", string(bytes))
	return strings.Replace(strings.TrimSpace(string(bytes)), "\n", "", -1)
}
