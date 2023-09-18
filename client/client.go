package client

// func NewClient(accessKey, secretKey string) *Client {
// 	return &Client{AccessKey: accessKey, SecretKey: secretKey}
// }

type Option func(*Client)

func UseGateway(url string) Option {
	return func(c *Client) {
		c.GatewayHost = url
		c.UseGateway = true
	}
}
func SetAkSk(ak, sk string) Option {
	return func(c *Client) {
		c.AccessKey = ak
		c.SecretKey = sk
	}
}

func NewClient(opts ...Option) *Client {
	client := &Client{}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

type Client struct {
	AccessKey   string
	SecretKey   string
	GatewayHost string
	UseGateway  bool
}

// NewClient(
// 	UseGateway("http://xapi-gateway.com:8080/api/name")
// 	UseGateway("http://localhost:8002")
// )
