package testx

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func Test(testEngine *gin.Engine, request *http.Request) (response *httptest.ResponseRecorder) {
	response = httptest.NewRecorder()
	testEngine.ServeHTTP(response, request)
	return
}
