package client

import (
	"net/url"
)

//----------------------------------
// https://api.aa1.cn/doc/myip.html
// 本地IP
// API简介：显示当前设备IP，欢迎对接
//----------------------------------

func (c *Client) Request_13(param, interfaceId string) ([]byte, error) {
	juheURL := "https://v.api.aa1.cn/api/myip/"

	params := url.Values{}

	params.Set("aa1", "json") //json  js  js-txt  text

	return c.Get(interfaceId, juheURL, params)
}
