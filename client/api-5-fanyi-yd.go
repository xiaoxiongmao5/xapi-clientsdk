package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/api-fanyi-yd.html
// 有道翻译
// API简介：在线中英互译
// Content-Type: text/html; charset=UTF-8
//----------------------------------

type Api_5_Param struct {
	Msg string `json:"msg"` //翻译内容
	// Type string `json:"type"` //请输入翻译类型（1代表中-英，2代表英-中，3代表中<=>英【自动检测翻译】）
}

func (c *Client) Api_5(param, transinfo3 string) ([]byte, error) {
	interfaceId := "5"
	juheURL := "https://v.api.aa1.cn/api/api-fanyi-yd/index.php"

	var requestParam Api_5_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("msg", strings.Trim(requestParam.Msg, " "))
	params.Set("type", "3")

	return c.Get(interfaceId, juheURL, params)
}
