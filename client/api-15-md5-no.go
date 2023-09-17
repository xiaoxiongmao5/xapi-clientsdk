package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/api-md5.html
// MD5在线加解密
// API简介：加解密MD5，仅简单的md5解密
//----------------------------------

type Api_15_Param struct {
	// Act string `json:"act"` //加密 / 解密(解密是坏的)
	Md5 string `json:"md5"` //输文字或加密后的代码
}

func (c *Client) Api_15(param, transinfo3 string) ([]byte, error) {
	interfaceId := "15"
	// juheURL := "https://v.api.aa1.cn/api/api-md5/go.php"
	juheURL := GATEWAY_HOST

	var requestParam Api_15_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("act", "加密")
	params.Set("md5", strings.Trim(requestParam.Md5, " "))

	return c.Get(interfaceId, juheURL, params)
}
