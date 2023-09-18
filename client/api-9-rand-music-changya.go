package client

import (
	"net/url"
)

//----------------------------------
// https://api.aa1.cn/doc/changya-xx.html
// 随机唱鸭歌曲
// API简介：随机返回一首唱鸭歌曲
// Content-Type: text/plain; charset=utf-8
//----------------------------------

func (c *Client) Api_9(param, transinfo3 string) ([]byte, error) {
	interfaceId := "9"
	//请求地址
	juheURL := "https://v2.api-m.com/api/changya"

	//初始化参数
	params := url.Values{}

	//发送请求
	return c.Get(interfaceId, juheURL, params)
}
