package testx

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestJson(testEngine *gin.Engine, method, url string, headers map[string]string, jsonData map[string]any) (response *httptest.ResponseRecorder) {
	// 将 map 序列化为 JSON
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// 创建一个 bytes.Reader 从 JSON 字节
	reader := bytes.NewReader(jsonBytes)

	// 构造请求
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		panic(err)
	}

	// 设置请求头
	request.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	// 发布模式
	gin.SetMode(gin.ReleaseMode)

	return Test(testEngine, request)
}
