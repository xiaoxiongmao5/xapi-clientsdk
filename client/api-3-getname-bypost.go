package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Api_3_Param struct {
	Name string `json:"name"`
}

// 使用POST方法像服务器发送USER对象，并获取服务器返回的结果
func (c *Client) Api_3(param, transinfo3 string) (bodyBytes []byte, err error) {
	interfaceId := "3"
	juheURL := ""
	if c.UseGateway {
		juheURL = c.GatewayHost
	}
	var requestParam Api_3_Param

	// 解析 JSON 字符串并填充实例
	err = json.Unmarshal([]byte(param), &requestParam)
	if err != nil {
		LogErr(interfaceId, "参数解析 JSON 失败", err)
		return
	}

	// // 构建查询字符串，将其附加到URL上
	// params := url.Values{}
	// params.Set("name", requestParam.Name)
	// if len(params) > 0 {
	// 	juheURL += "?" + params.Encode()
	// }

	body, err := json.Marshal(requestParam)
	if err != nil {
		LogErr(interfaceId, "Failed to Marshal", err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", juheURL, bytes.NewReader(body))
	if err != nil {
		LogErr(interfaceId, "Failed to create request", err)
		return
	}

	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", interfaceId)
	req.Header = headers
	req.Header.Add("Content-Type", "application-json")

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
