package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/text2img.html
// 文本转图片
// API简介：文本生成图片，文本转图片由平平免费API提供（api.setbug.com）
// Content-Type: image/png
//----------------------------------

type Api_16_Param struct {
	Text string `json:"text"`
}

func (c *Client) Api_16(param, transinfo3 string) ([]byte, error) {
	interfaceId := "16"
	juheURL := "http://api.setbug.com/tools/text2image/"

	var requestParam Api_16_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("text", strings.Trim(requestParam.Text, " "))

	return c.Get(interfaceId, juheURL, params)
}
