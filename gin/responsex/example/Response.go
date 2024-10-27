package example

import (
	"github.com/5-say/go-tool/gin/responsex"
	"github.com/gin-gonic/gin"
)

// Response ..
func Response(c *gin.Context) responsex.ResponseT {
	return responsex.ResponseT{
		C: c,

		DefaultSuccessHTTPCode: 200,
		DefaultSuccessMessage:  "success",

		DefaultErrorHTTPCode: 422,
		DefaultErrorMessage:  "error",

		SuccessDataWarp: func(messageFormat []any, data any) any {
			var message = Language.Get(messageFormat)

			var resource any = gin.H{}
			if data != nil {
				resource = data
			}

			return gin.H{
				"message":  message,
				"resource": resource,
			}
		},

		ErrorDataWarp: func(messageFormat []any, data any, err error) any {
			var message = Language.Get(messageFormat)
			if message == "error" && err != nil {
				message = err.Error()
			}

			var resource any = gin.H{}
			if data != nil {
				resource = data
			}

			return gin.H{
				"message":  message,
				"resource": resource,
			}
		},
	}
}
