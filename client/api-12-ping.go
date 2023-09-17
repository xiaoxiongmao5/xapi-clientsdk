package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/api-ping.html
// Ping测试
// API简介：一键ping网站延迟
//----------------------------------

type Api_12_Param struct {
	Url string `json:"url"` //测试的网址
}

func (c *Client) Api_12(param, transinfo3 string) ([]byte, error) {
	interfaceId := "12"
	// juheURL := "https://v.api.aa1.cn/api/api-ping/ping.php"
	juheURL := GATEWAY_HOST

	var requestParam Api_12_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("url", strings.Trim(requestParam.Url, " "))

	return c.Get(interfaceId, juheURL, params)
}
