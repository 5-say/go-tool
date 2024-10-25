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

		SuccessDataWarp: func(message string, data any) any {
			var resource any = gin.H{}
			if data != nil {
				resource = data
			}
			return gin.H{
				"message":  message,
				"resource": resource,
			}
		},

		ErrorDataWarp: func(message string, data any, err error) any {
			var resource any = gin.H{}
			if data != nil {
				resource = data
			}
			if message == "error" && err != nil {
				message = err.Error()
			}
			return gin.H{
				"message":  message,
				"resource": resource,
			}
		},
	}
}
