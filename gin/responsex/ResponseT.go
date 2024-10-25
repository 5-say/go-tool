package responsex

import (
	"github.com/gin-gonic/gin"
)

type ResponseT struct {
	C *gin.Context

	DefaultSuccessHTTPCode int
	DefaultSuccessMessage  string

	DefaultErrorHTTPCode int
	DefaultErrorMessage  string

	SuccessDataWarp func(message string, data any) any
	ErrorDataWarp   func(message string, data any, err error) any
}
