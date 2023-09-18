package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/api-port.html
// 查询域名whois
// API简介：域名whois查询，支持300+后缀
// Content-Type: text/html; charset=UTF-8
//----------------------------------

type Api_6_Param struct {
	Domain string `json:"domain"` //域名 https://v.api.aa1.cn/api/whois/index.php?domain=qq.com
}

func (c *Client) Api_6(param, transinfo3 string) ([]byte, error) {
	interfaceId := "6"
	juheURL := "https://v.api.aa1.cn/api/whois/index.php"

	var requestParam Api_6_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("domain", strings.Trim(requestParam.Domain, " "))

	return c.Get(interfaceId, juheURL, params)
}
