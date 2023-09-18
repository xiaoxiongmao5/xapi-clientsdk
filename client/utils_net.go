package client

import (
	"io"
	"net/http"
	"net/url"
)

// get 网络请求
func (c *Client) Get(interfaceId string, apiURL string, params url.Values) (rs []byte, err error) {
	if c.UseGateway {
		apiURL = c.GatewayHost
	}
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		LogErr(interfaceId, "解析url错误", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()

	// 创建一个具有自定义请求头的http.Client
	client := &http.Client{}

	// 创建一个HTTP请求对象并设置请求方法、URL和其他头部信息
	req, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		LogErr(interfaceId, "Failed to create request", err)
		return nil, err
	}
	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", interfaceId)
	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		LogErr(interfaceId, "Failed to make request", err)
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func Post(interfaceId string, apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
