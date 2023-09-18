package client

import (
	"net/url"
)

//----------------------------------
// https://api.aa1.cn/doc/yiyan.html
// 每日一言
// API简介：每日一言，欢迎对接
//----------------------------------

func (c *Client) Api_14(param, transinfo3 string) ([]byte, error) {
	interfaceId := "14"
	juheURL := "https://v.api.aa1.cn/api/yiyan/index.php"

	params := url.Values{}

	return c.Get(interfaceId, juheURL, params)
}
