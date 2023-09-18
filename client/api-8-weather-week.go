package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/api-tianqi-3.html
// 查询一周天气
// API简介：较稳天气接口，可对接
// Content-Type: text/html; charset=UTF-8
//----------------------------------

type Api_8_Param struct {
	Msg string `json:"msg"` //要查询的城市，如：温州、上海、北京
	// Type string `json:"type"` //查询的城市 - 区域级天气（如和平区）
}

// 1.根据城市查询天气
func (c *Client) Api_8(param, transinfo3 string) ([]byte, error) {
	interfaceId := "8"
	//请求地址
	juheURL := "https://v.api.aa1.cn/api/api-tianqi-3/index.php"

	var requestParam Api_8_Param
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
	params.Set("msg", strings.Trim(requestParam.Msg, " "))
	params.Set("type", "1")

	//发送请求
	return c.Get(interfaceId, juheURL, params)
}
