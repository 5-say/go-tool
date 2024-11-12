package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// 发起 POST 请求，发送 json 参数
func PostJSON(urlStr string, jsonParams any, headers map[string]string) ResponseT {
	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(jsonParams)
	if err != nil {
		return ResponseT{Error: errors.WithStack(err)}
	}

	// 创建一个新的 HTTP 请求对象
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		return ResponseT{Error: errors.WithStack(err)}
	}

	// 设置请求头信息
	req.Header.Set("Content-Type", "application/json") // 设置请求头信息，表明发送的数据是 JSON 格式
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 创建一个 HTTP 客户端，并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseT{Error: errors.WithStack(err)}
	}
	defer resp.Body.Close()

	// 读取响应信息
	body, err := io.ReadAll(resp.Body)
	return ResponseT{Error: errors.WithStack(err), Body: body, StatusCode: resp.StatusCode}
}
