package client

import (
	"net/url"
	"proxy/cons"
	"proxy/plugin/clash"
	"proxy/tools"
)

func CheckJob() error {

	delay, err := clash.GetProxyDelay("xg", url.PathEscape(cons.TestUrl), 10000)
	if err != nil {
		tools.Logger.Error("err : ", err)
	}

	if delay.Delay != 0 {
		tools.Logger.Info("test success , delay : ", delay.Delay)
		//return nil
	}

	url := "http://" + cons.ServerIP + ":" + cons.ServerHttpPort + "/server/incr-port/"

	var incrPort map[string]any
	err = tools.UnmarshalAuthRequest("get", url, nil, nil, &incrPort, cons.HttpUser, cons.HttpPassword)
	if err != nil {
		tools.Logger.Info("err : ", err)
		return nil
	}
	port := incrPort["data"].(map[string]any)["port"].(string)

	tools.Logger.Info("after incr port : ", port)

	return nil

}
