package testx

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestFormAuth(testEngine *gin.Engine, url, token string, queryParams map[string]any) (response *httptest.ResponseRecorder) {
	return TestForm(testEngine, url, map[string]string{
		"Authorization": token,
	}, queryParams)
}
