package clash

import (
	"encoding/json"
	"fmt"
	"proxy/cons"
	"proxy/tools"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`\[(.+?)\](.+?)lAddr=(.+?)rAddr=(.+?)mode=(.+?)rule=(.+?)proxy=(.+)`)

type LogMessage struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func (msg *LogMessage) ToLog() *Log {
	match := reg.FindAllStringSubmatch(msg.Payload, -1)
	if len(match) != 0 && len(match[0]) != 0 {
		l := &Log{
			Type:    msg.Type,
			Socket:  strings.TrimSpace(match[0][1]),
			Message: strings.TrimSpace(match[0][2]),
			LAddr:   strings.TrimSpace(match[0][3]),
			RAddr:   strings.TrimSpace(match[0][4]),
			Mode:    strings.TrimSpace(match[0][5]),
			Rule:    strings.TrimSpace(match[0][6]),
			Proxy:   strings.TrimSpace(match[0][7]),
		}
		return l
	}
	return &Log{Type: msg.Type, Error: msg.Payload}
}

type Log struct {
	Type    string `json:"type"`
	Socket  string `json:"socket"`
	Message string `json:"message"`
	LAddr   string `json:"local_addr"`
	RAddr   string `json:"remote_addr"`
	Mode    string `json:"mode"`
	Rule    string `json:"rule"`
	Proxy   string `json:"proxy"`
	Error   string `json:"error"`
}

type Traffic struct {
	Up   uint64 `json:"up"`
	Down uint64 `json:"down"`
}

type LogLevel string

const (
	LevelError   LogLevel = "error"
	LevelInfo    LogLevel = "info"
	LevelWarning LogLevel = "warning"
	LevelDebug   LogLevel = "debug"
)

func GetLogs(level LogLevel) (chan *Log, error) {
	logChan := make(chan *Log, 1024)

	headers := map[string]string{"level": string(level)}
	AddSecretHeader(&headers)
	url := cons.ClashServer + "/logs"
	resp, err := tools.Request("get", url, headers, nil)
	if err != nil {
		return logChan, err
	}

	tools.HandleStreamResp(resp, func(line []byte) (stop bool) {
		msg := &LogMessage{}
		if err := json.Unmarshal(line, msg); err != nil {
			return true
		}
		logChan <- msg.ToLog()
		return false
	})

	return logChan, nil
}

func GetTraffic(handler func(traffic *Traffic) (stop bool)) error {
	url := cons.ClashServer + "/traffic"
	headers := map[string]string{}
	AddSecretHeader(&headers)
	resp, err := tools.Request("get", url, headers, nil)
	if err != nil {
		return err
	}

	tools.HandleStreamResp(resp, func(line []byte) (stop bool) {
		traffic := &Traffic{}
		if err := json.Unmarshal(line, traffic); err != nil {
			return true
		}
		if _stop := handler(traffic); _stop {
			return true
		}
		return false
	})
	return nil
}

type Proxies struct {
	All     []string   `json:"all"`
	History []*History `json:"history"`
	Name    string     `json:"name"`
	Now     string     `json:"now"`
	Type    string     `json:"type"`
	UDP     bool       `json:"udp"`
}

type History struct {
	Time      string `json:"time"`
	Delay     int    `json:"delay"`
	MeanDelay int    `json:"meanDelay"`
}

type Proxy struct {
	History []*History `json:"history"`
	Name    string     `json:"name"`
	Type    string     `json:"type"`
	UDP     bool       `json:"udp"`
}

func GetProxies() (map[string]*Proxies, error) {
	container := struct {
		Proxies map[string]*Proxies `json:"proxies"`
	}{}
	url := cons.ClashServer + "/proxies"
	headers := map[string]string{}
	AddSecretHeader(&headers)
	err := tools.UnmarshalRequest("get", url, headers, nil, &container)
	if err != nil {
		return nil, err
	}
	return container.Proxies, nil
}

func GetProxyMessage(name string) (*Proxy, error) {
	proxy := &Proxy{}
	route := "/proxies/" + name
	url := cons.ClashServer + route
	headers := map[string]string{}
	AddSecretHeader(&headers)
	err := tools.UnmarshalRequest("get", url, headers, nil, &proxy)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

type ProxyDelay struct {
	Delay     int    `json:"delay"`
	MeanDelay int    `json:"meanDelay"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}

func GetProxyDelay(name string, url string, timeout int) (*ProxyDelay, error) {
	proxyDelay := &ProxyDelay{}
	route := fmt.Sprintf("/proxies/%s/delay?url=%s&timeout=%d", name, url, timeout)
	url2 := cons.ClashServer + route
	headers := map[string]string{}
	AddSecretHeader(&headers)
	err := tools.UnmarshalRequest("get", url2, headers, nil, &proxyDelay)
	if err != nil {
		return nil, err
	}
	return proxyDelay, nil
}

func SwitchProxy(selector, name string) error {
	route := "/proxies/" + selector
	headers := map[string]string{"Content-Type": "application/json"}
	AddSecretHeader(&headers)
	body := map[string]interface{}{"name": name}
	url := cons.ClashServer + route
	code, content, err := tools.EasyRequest("put", url, headers, body)
	if err != nil {
		return err
	}

	switch code {
	case 204:
		return nil
	case 400, 404:
		return fmt.Errorf(string(content))
	default:
		return fmt.Errorf("unknown error: %s", string(content))
	}
}

type Config struct {
	Port           int      `json:"port"`
	SocksPort      int      `json:"socks-port"`
	RedirPort      int      `json:"redir-port"`
	TproxyPort     int      `json:"tproxy-port"`
	MixedPort      int      `json:"mixed-port"`
	Authentication []string `json:"authentication"`
	AllowLan       bool     `json:"allow-lan"`
	BindAddress    string   `json:"bind-address"`
	Mode           string   `json:"mode"`
	LogLevel       string   `json:"log-level"`
	IPV6           bool     `json:"ipv6"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	url := cons.ClashServer + "/configs"
	headers := map[string]string{}
	AddSecretHeader(&headers)
	err := tools.UnmarshalRequest("get", url, headers, nil, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type Rule struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Proxy   string `json:"proxy"`
}

func GetRules() ([]*Rule, error) {
	container := struct {
		Rules []*Rule `json:"rules"`
	}{}
	url := cons.ClashServer + "/rules"
	headers := map[string]string{}
	AddSecretHeader(&headers)
	err := tools.UnmarshalRequest("get", url, headers, nil, &container)
	if err != nil {
		return nil, err
	}
	return container.Rules, nil
}

// EnableConfig 这个接口不会影响 external-controller 和 secret 的值
func EnableConfig(path string) error {
	headers := map[string]string{"Content-Type": "application/json"}
	AddSecretHeader(&headers)
	body := map[string]interface{}{"path": path}

	url := cons.ClashServer + "/configs"
	code, content, err := tools.EasyRequest("put", url, headers, body)
	if err != nil {
		return err
	}
	if code != 200 {
		return fmt.Errorf("unknown error: %s", string(content))
	}
	return nil
}

func SetConfig(port, socksPort int, redirPort string, allowLan bool, mode, logLevel string) error {
	headers := map[string]string{"Content-Type": "application/json"}
	AddSecretHeader(&headers)
	body := map[string]interface{}{
		"port":       port,
		"socks-port": socksPort,
		"redir-port": redirPort,
		"allow-lan":  allowLan,
		"mode":       mode,
		"log-level":  logLevel,
	}
	url := cons.ClashServer + "/configs"
	code, content, err := tools.EasyRequest("patch", url, headers, body)
	if err != nil {
		return err
	}
	if code != 200 {
		return fmt.Errorf("unknown error: %s", string(content))
	}
	return nil
}

func AddSecretHeader(headers *map[string]string) {
	value := fmt.Sprintf("Bearer %s", cons.ClashSecret)
	(*headers)["Authorization"] = value
}
