package client

// 网关地址
// var GATEWAY_HOST = "http://localhost:8002"

var GATEWAY_HOST = "http://xapi-gateway.com:8080/api/name"

func NewClient(accessKey, secretKey string) *Client {
	return &Client{AccessKey: accessKey, SecretKey: secretKey}
}

type Client struct {
	AccessKey string
	SecretKey string
}
