package client

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
// 		LogErr(funcName, "Failed to create request", err)
// 		return
// 	}

// 	// 构建请求头
// 	headers := getRequestHeaders(c.AccessKey, c.SecretKey, "")
// 	req.Header = headers

// 	response, err := client.Do(req)
// 	if err != nil {
// 		LogErr(funcName, "Failed to make request", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// 解析响应
// 	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
// 		return
// 	}

// 	return data, nil
// }
