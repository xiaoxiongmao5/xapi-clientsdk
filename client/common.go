package client

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func LogErr(funcName, desc string, err error) {
	if err == nil {
		fmt.Printf("[%s]: %s", funcName, desc)
		return
	}
	fmt.Printf("[%s]: %s。 err=%s", funcName, desc, err.Error())
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
func getRequestHeaders(accessKey, secretkey, requestBody, interfaceId string) http.Header {
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
	headers.Set("interfaceId", interfaceId)

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
