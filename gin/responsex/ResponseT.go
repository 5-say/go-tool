package responsex

import (
	"github.com/gin-gonic/gin"
)

type ResponseT struct {
	C *gin.Context

	DefaultSuccessHTTPCode int
	DefaultErrorHTTPCode   int

	SuccessDataWarp func(data interface{}) interface{}
	ErrorDataWarp   func(data interface{}) interface{}
}
