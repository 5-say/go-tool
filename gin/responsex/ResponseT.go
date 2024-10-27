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

	SuccessDataWarp func(messageFormat []any, data any) any
	ErrorDataWarp   func(messageFormat []any, data any, err error) any
}
