package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Api_2_Param struct {
	Name string `json:"name"`
}

func (c *Client) Api_2(param, transinfo3 string) (bodyBytes []byte, err error) {
	interfaceId := "2"
	juheURL := ""
	if c.UseGateway {
		juheURL = c.GatewayHost
	}
	var requestParam Api_2_Param

	// 解析 JSON 字符串并填充实例
	err = json.Unmarshal([]byte(param), &requestParam)
	if err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return
	}

	// 构建查询字符串，将其附加到URL上
	params := url.Values{}

	params.Set("name", requestParam.Name)
	var Url *url.URL
	Url, _ = url.Parse(juheURL)
	Url.RawQuery = params.Encode()

	// body, err := json.Marshal(requestParam)
	// if err != nil {
	// 	LogErr(interfaceId, "Failed to Marshal", err)
	// 	return
	// }

	client := &http.Client{}

	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		LogErr(interfaceId, "Failed to create request", err)
		return
	}

	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", interfaceId)
	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		LogErr(interfaceId, "Failed to make request", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		LogErr(interfaceId, "Failed to read response", err)
		return
	}

	// statusCode = response.StatusCode
	// contentType = response.Header.Get("Content-Type")
	return
}
