package client

import (
	"net/url"
)

//----------------------------------
// https://api.aa1.cn/doc/uuid.html
// UUID生成
// API简介：随机生成uuid
// Content-Type: text/json;charset=UTF-8
//----------------------------------

func (c *Client) Api_17(param, transinfo3 string) ([]byte, error) {
	interfaceId := "17"
	juheURL := "https://v.api.aa1.cn/api/uuid/"

	params := url.Values{}

	return c.Get(interfaceId, juheURL, params)
}
