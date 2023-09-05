package client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 网关地址
// var GATEWAY_HOST = "http://localhost:8002/api/name"

var GATEWAY_HOST = "http://localhost:8080/api/name"

type Client struct {
	AccessKey string
	SecretKey string
}

func NewClient(accessKey, secretKey string) *Client {
	return &Client{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

// 计算API签名
func calculateSignature(accessKey, secretKey, nonce, timestamp, requestBody string) string {
	// 将参数拼接成一个字符串
	concatenatedString := accessKey + nonce + timestamp + requestBody + secretKey

	// 计算 MD5 值
	signature := md5.Sum([]byte(concatenatedString))
	return hex.EncodeToString(signature[:])
}

// 获得请求头
func getRequestHeaders(accessKey, secretkey, requestBody, gateway_transdata string) http.Header {
	headers := make(http.Header)

	// 生成 nonce : 一个包含100个随机数字的字符串
	nonce := GenetateRandomString(100)

	// 当前时间戳（秒级别）
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 计算签名
	signature := calculateSignature(accessKey, secretkey, nonce, timestamp, requestBody)

	// 设置请求头
	headers.Set("accessKey", accessKey)
	headers.Set("nonce", nonce)
	headers.Set("timestamp", timestamp)
	headers.Set("sign", signature)
	headers.Set("gateway_transdata", gateway_transdata)

	return headers
}

/** 生成包含N个随机数字的字符串
 */
func GenetateRandomString(length int) string {
	// 设置随机数种子，以确保每次运行生成的随机数都不同
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 定义一个包含数字字符的字符集
	charset := "0123456789"
	charsetLength := len(charset)

	// 生成随机数字并拼接字符串
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(charsetLength)
		randomChar := charset[randomIndex]
		randomString[i] = randomChar
	}
	return string(randomString)
}

type GetNameByGetParam struct {
	Name string `json:"name"`
}

func (c *Client) GetNameByGet(param, gateway_transdata string) (statusCode int, contentType string, bodyBytes []byte, err error) {
	var requestParam GetNameByGetParam

	// 解析 JSON 字符串并填充实例
	err = json.Unmarshal([]byte(param), &requestParam)
	if err != nil {
		fmt.Println("参数解析 JSON 失败:", err)
		return
	}

	// 构建查询字符串，将其附加到URL上
	params := url.Values{}
	params.Set("name", requestParam.Name)

	// 构建包含查询参数的URL
	fullURL := GATEWAY_HOST
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	client := &http.Client{}
	method := "GET"

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		fmt.Println("GetNameByGet: Failed to create request")
		return
	}

	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", gateway_transdata)
	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("GetNameByGet: Failed to make request")
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("GetNameByPost: Failed to read response, err=", err)
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
	var requestParam GetNameByPostParam

	// 解析 JSON 字符串并填充实例
	err = json.Unmarshal([]byte(param), &requestParam)
	if err != nil {
		fmt.Println("参数解析 JSON 失败:", err)
		return
	}

	// 构建查询字符串，将其附加到URL上
	params := url.Values{}
	params.Set("name", requestParam.Name)

	// 构建包含查询参数的URL
	fullURL := GATEWAY_HOST
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	body, err := json.Marshal(requestParam)
	if err != nil {
		fmt.Println("GetNameByPost: Failed to Marshal")
		return
	}

	client := &http.Client{}
	method := "POST"

	req, err := http.NewRequest(method, fullURL, bytes.NewReader(body))
	if err != nil {
		fmt.Println("GetNameByPost: Failed to create request")
		return
	}

	// 构建请求头
	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "", gateway_transdata)
	req.Header = headers
	req.Header.Add("Content-Type", "application-json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("GetNameByPost: Failed to make request")
		return
	}
	defer response.Body.Close()

	// 读取响应体，将响应体内容原封不动地返回给前端
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("GetNameByPost: Failed to read response, err=", err)
		return
	}

	statusCode = response.StatusCode
	contentType = response.Header.Get("Content-Type")
	// todo 判断 contentType 是否和注册接口信息要求一致

	return
}

/*
*
这种写法的请求方，使用反射，处理返回值逻辑：

	// 如果没有返回值或提取值无效，返回错误
	if len(result) < 2 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No valid data found",
		})
		return
	}

	fmt.Printf("result=%v, len=%v, type=%T \n", result, len(result), result)

	// 提取第一个返回值 对应 返回值 data interface{}
	firstResult := result[0]

	// 提取第二个返回值 对应 返回值 error
	bodyErrorValue := result[1]
	bodyError, ok := bodyErrorValue.Interface().(error)
	if ok {
		// 如果存在error
		fmt.Printf("bodyError=%v type=%T\n", bodyError, bodyError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "调用接口返回失败",
		})
		return
	}

	// 检查提取的值是否有效
	if firstResult.IsValid() {
		// 尝试将其转换为 interface{} 类型
		data := firstResult.Interface()

		// // 将 data 编码为 JSON 格式
		// jsonData, err := json.Marshal(data)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": "Failed to marshal data",
		// 	})
		// 	return
		// }

		// 返回 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
		return
	}
*/
// func (c *Client) GetNameByGet_old(name string) (data interface{}, err error) {

// 	// 构建查询字符串，将其附加到URL上
// 	params := url.Values{}
// 	params.Set("name", name)

// 	// 构建包含查询参数的URL
// 	fullURL := GATEWAY_HOST
// 	if len(params) > 0 {
// 		fullURL += "?" + params.Encode()
// 	}

// 	client := &http.Client{}
// 	method := "GET"

// 	req, err := http.NewRequest(method, fullURL, nil)
// 	if err != nil {
// 		fmt.Println("GetNameByGet: Failed to create request")
// 		return
// 	}

// 	// 构建请求头
// 	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "")
// 	req.Header = headers

// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("GetNameByGet: Failed to make request")
// 		return
// 	}
// 	defer response.Body.Close()

// 	// 解析响应
// 	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
// 		return
// 	}

// 	return data, nil
// }
