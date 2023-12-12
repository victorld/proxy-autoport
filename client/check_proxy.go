package client

import (
	"net/url"
	"os"
	"os/exec"
	"proxy/cons"
	"proxy/plugin/clash"
	"proxy/tools"
)

func CheckJob() (string, error) {

	delay, err := clash.GetProxyDelay("xg", url.PathEscape(cons.TestUrl), 10000)
	if err != nil {
		tools.Logger.Error("get proxy delay error : ", err)
		return "get proxy delay error", err
	}

	if delay.Delay != 0 {
		incrPortUrl := "http://" + cons.ServerIP + ":" + cons.ServerHttpPort + "/server/get-port/"
		var getPort map[string]any
		err = tools.UnmarshalAuthRequest("get", incrPortUrl, nil, nil, &getPort, cons.HttpUser, cons.HttpPassword)
		if err != nil {
			tools.Logger.Error("incr port error : ", err)
			return "incr port error", err
		}
		port := getPort["data"].(map[string]any)["port"].(string)
		tools.Logger.Info("test success , delay : ", delay.Delay, " , port : ", port)
		return "test success , port no change", nil
	}

	incrPortUrl := "http://" + cons.ServerIP + ":" + cons.ServerHttpPort + "/server/incr-port/"
	var incrPort map[string]any
	err = tools.UnmarshalAuthRequest("get", incrPortUrl, nil, nil, &incrPort, cons.HttpUser, cons.HttpPassword)
	if err != nil {
		tools.Logger.Error("incr port error : ", err)
		return "incr port error", err
	}
	port := incrPort["data"].(map[string]any)["port"].(string)

	tools.Logger.Info("after incr port : ", port)

	err = os.Chmod(cons.ClientChangePortShellPath, 0777)
	if err != nil {
		tools.Logger.Error("chmod config error : ", err)
	}

	command := exec.Command("/bin/sh", "-c", cons.ClientChangePortShellPath+" "+port+" "+cons.ClientLocalConfitFilePath) //初始化Cmd
	bytes, err := command.Output()
	if err != nil {
		tools.Logger.Error("client change port error : ", err)
		return "client change port errpr", err
	}
	tools.Logger.Info("shell out : ", string(bytes))

	err = clash.EnableConfig(cons.ClashServerConfitFilePath)
	if err != nil {
		tools.Logger.Error("client update config errpr : ", err)
		return "client update config errpr", err
	}

	tools.Logger.Info("update port success !")
	return "update port success !", nil

}
