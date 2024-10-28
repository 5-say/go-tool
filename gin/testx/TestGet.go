package testx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func TestGet(testEngine *gin.Engine, urlStr string, headers map[string]string, queryParams map[string]any) (response *httptest.ResponseRecorder) {
	// 构建带有查询参数的 URL
	var params = url.Values{}
	for k, v := range queryParams {
		params.Add(k, fmt.Sprintf("%v", v))
	}
	var (
		query = params.Encode()
		url   = urlStr + "?" + query
	)

	// 构造请求
	request, err := http.NewRequest("GET", url, nil)
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
