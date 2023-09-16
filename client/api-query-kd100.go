package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/kd100.html
// 快递100快递查询
// API简介：调用快递100快递接口查询
//----------------------------------

type Request_11_Param struct {
	Dh string `json:"dh"`
}

func (c *Client) Request_11(param, interfaceId string) ([]byte, error) {
	juheURL := "https://free.wqwlkj.cn/wqwlapi/kd100.php"

	var requestParam Request_11_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("dh", strings.Trim(requestParam.Dh, " "))

	return c.Get(interfaceId, juheURL, params)
}
