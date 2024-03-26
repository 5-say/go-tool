package testx

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestJsonAuth(testEngine *gin.Engine, method, url, token string, jsonData map[string]any) (response *httptest.ResponseRecorder) {
	return TestJson(testEngine, method, url, map[string]string{
		"Authorization": token,
	}, jsonData)
}
