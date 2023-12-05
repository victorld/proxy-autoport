package server

import (
	"fmt"
	"os/exec"
	"proxy/tools"
	"strconv"
	"strings"
)

func ChangePort() {
	content := tools.ReadFileLine(tools.LineNumber, tools.ConfigFilePath)
	split := strings.TrimSpace(strings.Split(strings.Split(content, ":")[1], ",")[0])
	port, _ := strconv.Atoi(split)
	if port == 0 {
		port = tools.DefaultPort
	}
	newPort := port + 1
	fmt.Println("newPort : ", newPort)
	command := exec.Command("/bin/sh", "-c", tools.ShellPath+" "+strconv.Itoa(newPort)) //初始化Cmd
	bytes, err := command.Output()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}

	fmt.Println(string(bytes))
}