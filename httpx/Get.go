package httpx

import (
	"io"
	"net/http"
	"net/url"
)

// 发起 GET 请求
//
//	切记！`urlStr` 一定要带 `http://` 或 `http://` 前缀
//	否则 `client.Do(req)` 会抛出错误 `unsupported protocol scheme ""`
//
//	urlStr      string
//	queryParams map[string]string
//	headers     map[string]string
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
//	statusCode, err := httpx.Get(urlStr, queryParams, headers).BindJSON(&apiResult)
func Get(urlStr string, queryParams, headers map[string]string) ResponseT {
	// 构建带有查询参数的 URL
	var params = url.Values{}
	for k, v := range queryParams {
		params.Add(k, v)
	}
	var (
		query = params.Encode()
		url   = urlStr + "?" + query
	)

	// 创建一个新的 HTTP 请求对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseT{Error: err}
	}

	// 设置请求头信息
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
