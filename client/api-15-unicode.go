package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/unicode.html
// UniCode转码
// API简介：中文和uniCode互转
// Content-Type: text/html; charset=UTF-8
//----------------------------------

type Api_15_Param struct {
	Msg  string `json:"msg"`  //待解码/编码转换文本
	Type string `json:"type"` //bm：编码，jm：解码
}

func (c *Client) Api_15(param, transinfo3 string) ([]byte, error) {
	interfaceId := "15"
	juheURL := "https://zj.v.api.aa1.cn/api/unicode/"

	var requestParam Api_15_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("msg", strings.Trim(requestParam.Msg, " "))
	params.Set("type", strings.Trim(requestParam.Type, " "))

	return c.Get(interfaceId, juheURL, params)
}
