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

		SuccessDataWarp: func(message string, data any) any {
			var (
				dataObj any = gin.H{
					"message": message,
				}
				errorObj any = gin.H{}
			)
			if data != nil {
				dataObj = data
			}
			return gin.H{
				"data":  dataObj,
				"error": errorObj,
			}
		},

		ErrorDataWarp: func(message string, data any, err error) any {
			if message == "error" && err != nil {
				message = err.Error()
			}
			var (
				dataObj  any = gin.H{}
				errorObj any = gin.H{
					"code":    "",
					"message": message,
				}
			)
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
