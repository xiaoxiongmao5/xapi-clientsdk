package client

import (
	"net/url"
)

//----------------------------------
// https://api.aa1.cn/doc/times.html
// 输出当期时间戳
// API简介：时间戳输出
// Content-Type: text/html; charset=UTF-8
//----------------------------------

func (c *Client) Api_18(param, transinfo3 string) ([]byte, error) {
	interfaceId := "18"
	juheURL := "https://v.api.aa1.cn/api/times/"

	params := url.Values{}

	return c.Get(interfaceId, juheURL, params)
}
