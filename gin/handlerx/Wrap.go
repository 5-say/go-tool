package handlerx

import (
	"github.com/gin-gonic/gin"
)

// InternalHandler ..
type InternalHandler = func(c *gin.Context, r Response) *ResponseBag

// Wrap ..
//
// e.g.
//
//	handlerx.Wrap(c, func(c *gin.Context, r handlerx.Response) *handlerx.ResponseBag {
//		return r.Success()
//	})
func Wrap(c *gin.Context, internalHandler InternalHandler) {
	var r = internalHandler(c, Response{})
	// 构造响应
	if r != nil {
		c.AbortWithStatusJSON(r.httpStatusCode, map[string]any{
			"mark":     r.mark,
			"message":  r.message,
			"resource": r.resource,
		})
	}
}
