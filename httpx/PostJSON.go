package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// 发起 POST 请求，发送 json 参数
//
//	切记！`urlStr` 一定要带 `http://` 或 `http://` 前缀
//	否则 `client.Do(req)` 会抛出错误 `unsupported protocol scheme ""`
//
//	urlStr     string
//	jsonParams any
//	headers    map[string]string
//
//	return
//
//	httpx.ResponseT struct {
//		Error      error
//		Body       []byte
//		StatusCode int
//		BindJSON   func (obj any) statusCode int, err error // 解析响应，绑定到对象
//	}
//
// e.g.
//
//	statusCode, err := httpx.PostJSON(urlStr, jsonParams, headers).BindJSON(&apiResult)
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
