package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/ipquery.html
// IP归属地查询
// API简介：ip查询暂不支持域名查询
// Content-Type: application/json
//----------------------------------

type Api_12_Param struct {
	Ip string `json:"ip"` //IP
}

func (c *Client) Api_12(param, transinfo3 string) ([]byte, error) {
	interfaceId := "12"
	juheURL := "http://www.lpv4.cn:10000/api/ip"

	var requestParam Api_12_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("ip", strings.Trim(requestParam.Ip, " "))

	return c.Get(interfaceId, juheURL, params)
}
