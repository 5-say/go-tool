package handlerx

import (
	"github.com/gin-gonic/gin"
)

// InternalHandler ..
type InternalHandler = func(c *gin.Context, response Response) *ResponseBag

// Wrap ..
//
// ex:
//
//	handlerx.Wrap(c, func(c *gin.Context, response handlerx.Response) *handlerx.ResponseBag {
//		return response.Success()
//	})
func Wrap(c *gin.Context, internalHandler InternalHandler) {
	var r = internalHandler(c, Response{})
	// 构造响应
	if r != nil {
		c.AbortWithStatusJSON(r.httpStatusCode, map[string]any{
			"mark":    r.mark,
			"data":    r.data,
			"message": r.message,
		})
	}
}
