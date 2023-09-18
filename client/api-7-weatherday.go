package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/weathertq.html
// 查询当前天气
// API简介：php 天气查询
// Content-Type: text/html;charset=utf-8
//----------------------------------

type Api_7_Param struct {
	City string `json:"city"` //要查询的城市，如：温州、上海、北京
}

// 1.根据城市查询天气
func (c *Client) Api_7(param, transinfo3 string) ([]byte, error) {
	interfaceId := "7"
	//请求地址
	juheURL := "http://www.lpv4.cn:10000/api/weather"

	var requestParam Api_7_Param
	// 解析 JSON 字符串并填充实例
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	// var transinfo3Param Api_8_Transinfo3
	// if err := json.Unmarshal([]byte(param), &transinfo3Param); err != nil {
	// 	LogErr(interfaceId, "transinfo3参数解析 JSON 失败", err)
	// 	return nil, err
	// }

	//初始化参数
	params := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	params.Set("city", strings.Trim(requestParam.City, " "))

	//发送请求
	return c.Get(interfaceId, juheURL, params)
}
