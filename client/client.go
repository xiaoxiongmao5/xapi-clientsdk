package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// 网关地址
// var GATEWAY_HOST = "http://localhost:8002"

var GATEWAY_HOST = "http://xapi-gateway.com:8080"

func NewClient(accessKey, secretKey string) *Client {
	return &Client{AccessKey: accessKey, SecretKey: secretKey}
}

type Client struct {
	AccessKey string
	SecretKey string
}

type GetNameByGetParam struct {
	Name string `json:"name"`
}

func (c *Client) GetNameByGet(param, gateway_transdata string) (statusCode int, contentType string, bodyBytes []byte, err error) {
	funcName := "GetNameByGet"
	method, apiURL := "GET", "/api/name"
	fullURL := GATEWAY_HOST + apiURL
	var requestParam GetNameByGetParam

	// 解析 JSON 字符串并填充实例
	err = json.Unmarshal([]byte(param), &requestParam)
	if err != nil {
		LogErr(funcName, "参数解析 JSON 失败", err)
		return
	}

	// 构建查询字符串，将其附加到URL上
	params := url.Values{}
	params.Set("name", requestParam.Name)
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	// body, err := json.Marshal(requestParam)
	// if err != nil {
	// 	LogErr(funcName, "Failed to Marshal", err)
	// 	return
	// }

	client := &http.Client{}

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		LogErr(funcName, "Failed to create request", err)
		return
	}

	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", gateway_transdata)
	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		LogErr(funcName, "Failed to make request", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		LogErr(funcName, "Failed to read response", err)
		return
	}

	statusCode = response.StatusCode
	contentType = response.Header.Get("Content-Type")
	return
}

type GetNameByPostParam struct {
	Name string `json:"name"`
}

// 使用POST方法像服务器发送USER对象，并获取服务器返回的结果
func (c *Client) GetNameByPost(param, gateway_transdata string) (statusCode int, contentType string, bodyBytes []byte, err error) {
	funcName := "GetNameByPost"
	method, apiURL := "POST", "/api/name"
	fullURL := GATEWAY_HOST + apiURL
	var requestParam GetNameByPostParam

	// 解析 JSON 字符串并填充实例
	err = json.Unmarshal([]byte(param), &requestParam)
	if err != nil {
		LogErr(funcName, "参数解析 JSON 失败", err)
		return
	}

	// // 构建查询字符串，将其附加到URL上
	// params := url.Values{}
	// params.Set("name", requestParam.Name)
	// if len(params) > 0 {
	// 	fullURL += "?" + params.Encode()
	// }

	body, err := json.Marshal(requestParam)
	if err != nil {
		LogErr(funcName, "Failed to Marshal", err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, fullURL, bytes.NewReader(body))
	if err != nil {
		LogErr(funcName, "Failed to create request", err)
		return
	}

	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", gateway_transdata)
	req.Header = headers
	req.Header.Add("Content-Type", "application-json")

	response, err := client.Do(req)
	if err != nil {
		LogErr(funcName, "Failed to make request", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		LogErr(funcName, "Failed to read response", err)
		return
	}

	statusCode = response.StatusCode
	contentType = response.Header.Get("Content-Type")
	return
}
