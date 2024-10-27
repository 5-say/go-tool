package example

import (
	"github.com/5-say/go-tool/gin/responsex"
	"github.com/gin-gonic/gin"
)

// Response2 ..
func Response2(c *gin.Context) responsex.ResponseT {
	return responsex.ResponseT{
		C: c,

		DefaultSuccessHTTPCode: 200,
		DefaultSuccessMessage:  "success",

		DefaultErrorHTTPCode: 422,
		DefaultErrorMessage:  "error",

		SuccessDataWarp: func(messageFormat []any, data any) any {
			var message = Language.Get(messageFormat)

			var dataObj any = gin.H{
				"message": message,
			}
			if data != nil {
				dataObj = data
			}

			var errorObj any = gin.H{}

			return gin.H{
				"data":  dataObj,
				"error": errorObj,
			}
		},

		ErrorDataWarp: func(messageFormat []any, data any, err error) any {
			var message = Language.Get(messageFormat)
			if message == "error" && err != nil {
				message = err.Error()
			}

			var dataObj any = gin.H{}

			var errorObj any = gin.H{
				"code":    "",
				"message": message,
			}
			if data != nil {
				errorObj = data
			}

			return gin.H{
				"data":  dataObj,
				"error": errorObj,
			}
		},
	}
}
