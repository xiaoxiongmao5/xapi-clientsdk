package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://free-api.com/doc/69
// 天气预报
// 查询天气情况：文档、湿度、AQI、天气、风向等
// 天气预报调用示例代码 － 聚合数据
// 在线接口文档：http://www.juhe.cn/docs/73
//----------------------------------

type Api_8_Param struct {
	Cityname string `json:"cityname"` //要查询的城市，如：温州、上海、北京
}
type Api_8_Transinfo3 struct {
	APPKEY string `json:"appkey"` //应用APPKEY(应用详细页查询)
}

// 1.根据城市查询天气
func (c *Client) Api_8(param, transinfo3 string) ([]byte, error) {
	interfaceId := "8"
	//请求地址
	// juheURL := "http://op.juhe.cn/onebox/weather/query"
	juheURL := GATEWAY_HOST
	const APPKEY = "251518e073ef6c3c9504dd286c3f6a86" //您申请的APPKEY

	var requestParam Api_8_Param
	// 解析 JSON 字符串并填充实例
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	var transinfo3Param Api_8_Transinfo3
	if err := json.Unmarshal([]byte(param), &transinfo3Param); err != nil {
		LogErr(interfaceId, "transinfo3参数解析 JSON 失败", err)
		return nil, err
	}

	//初始化参数
	params := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	params.Set("cityname", strings.Trim(requestParam.Cityname, " "))
	params.Set("key", APPKEY)
	// params.Set("key", strings.Trim(transinfo3Param.APPKEY, " "))
	params.Set("dtype", "") //返回数据的格式,xml或json，默认json

	//发送请求
	return c.Get(interfaceId, juheURL, params)
	// data, err := c.Get(interfaceId, juheURL, params)
	// if err != nil {
	// 	fmt.Errorf("请求失败,错误信息:%v", err)
	// } else {
	// 	var netReturn map[string]interface{}
	// 	json.Unmarshal(data, &netReturn)
	// 	if netReturn["error_code"].(float64) == 0 {
	// 		fmt.Printf("接口返回result字段是:%v", netReturn["result"])
	// 	}
	// }
}
