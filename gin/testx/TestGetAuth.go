package testx

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestGetAuth(testEngine *gin.Engine, url, token string, queryParams map[string]any) (response *httptest.ResponseRecorder) {
	return TestGet(testEngine, url, map[string]string{
		"Authorization": token,
	}, queryParams)
}
