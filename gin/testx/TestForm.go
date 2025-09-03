package testx

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func TestForm(testEngine *gin.Engine, targetURL string, headers map[string]string, formData map[string]any) (response *httptest.ResponseRecorder) {
	// 准备表单数据
	formValues := url.Values{}
	for k, v := range formData {
		formValues.Set(k, v.(string))
	}

	// 创建请求 (Method, URL, Body)
	request, _ := http.NewRequest("POST", targetURL, strings.NewReader(formValues.Encode()))

	// 设置请求头
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	// 发布模式
	gin.SetMode(gin.ReleaseMode)

	return Test(testEngine, request)
}
