package httpx

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// 发起 POST 请求，发送 表单 参数
func Post(urlStr string, formParams, headers map[string]string) ResponseT {
	// 准备要发送的表单数据
	var params = url.Values{}
	for k, v := range formParams {
		params.Add(k, v)
	}

	// 将表单数据编码为请求体
	reqBody := strings.NewReader(params.Encode())

	// 创建一个新的 HTTP 请求对象
	req, err := http.NewRequest("POST", urlStr, reqBody)
	if err != nil {
		return ResponseT{Error: err}
	}

	// 设置请求头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // 设置请求头信息，表明发送的数据是表单格式
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 创建一个 HTTP 客户端，并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseT{Error: err}
	}
	defer resp.Body.Close()

	// 读取响应信息
	body, err := io.ReadAll(resp.Body)
	return ResponseT{Error: err, Body: body, StatusCode: resp.StatusCode}
}
