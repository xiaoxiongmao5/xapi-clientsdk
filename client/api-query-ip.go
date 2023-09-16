package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/ip-taobao.html
// IP归属地淘宝版
// API简介：IP归属地查询，数据来源于淘宝
//----------------------------------

type Request_10_Param struct {
	Ip string `json:"ip"`
}

func (c *Client) Request_10(param, interfaceId string) ([]byte, error) {
	juheURL := "https://zj.v.api.aa1.cn/api/ip-taobao/"

	var requestParam Request_10_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("ip", strings.Trim(requestParam.Ip, " "))

	return c.Get(interfaceId, juheURL, params)
}
