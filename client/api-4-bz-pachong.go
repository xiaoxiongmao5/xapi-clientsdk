package client

import (
	"encoding/json"
	"net/url"
	"strings"
)

//----------------------------------
// https://api.aa1.cn/doc/api-bz-scrapy.html
// 扒站接口爬虫版
// API简介：爬取目标站静态资源，接口由浑欲不胜簪维护
//----------------------------------

type Api_4_Param struct {
	Url string `json:"url"` //爬取的url
}

func (c *Client) Api_4(param, transinfo3 string) ([]byte, error) {
	interfaceId := "4"
	juheURL := "https://v.api.aa1.cn/api/api-bz/temp.php"

	var requestParam Api_4_Param
	if err := json.Unmarshal([]byte(param), &requestParam); err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return nil, err
	}

	params := url.Values{}

	params.Set("url", strings.Trim(requestParam.Url, " "))

	return c.Get(interfaceId, juheURL, params)
}
