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

type Api_11_Param struct {
	Dh string `json:"dh"`
}

func (c *Client) Api_11(param, transinfo3 string) ([]byte, error) {
	interfaceId := "11"
	// juheURL := "https://free.wqwlkj.cn/wqwlapi/kd100.php"
	juheURL := GATEWAY_HOST

	var requestParam Api_11_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("dh", strings.Trim(requestParam.Dh, " "))

	return c.Get(interfaceId, juheURL, params)
}
