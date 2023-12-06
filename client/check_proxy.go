package client

import (
	"net/url"
	"proxy/cons"
	"proxy/plugin/clash"
	"proxy/tools"
)

func CheckJob() {

	delay, err := clash.GetProxyDelay("xg", url.PathEscape(cons.TestUrl), 10000)
	if err != nil {
		tools.Logger.Error("err : ", err)
	}
	if delay.Delay != 0 {
		tools.Logger.Info("test success , delay : ", delay.Delay)
		return
	}

}
