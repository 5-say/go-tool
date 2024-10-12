package responsex

import "github.com/gin-gonic/gin"

// Success ..
//
// e.g.
//
//	r.Success()
//	r.Success(data interface{})
func (s ResponseT) Success(data ...interface{}) LogT {
	return s.SuccessCode(s.DefaultSuccessHTTPCode, data...)
}

// SuccessCode ..
//
// e.g.
//
//	r.SuccessCode(httpStatusCode int)
//	r.SuccessCode(httpStatusCode int, data interface{})
func (s ResponseT) SuccessCode(httpStatusCode int, data ...interface{}) LogT {
	// 构造响应数据
	var responseData interface{}
	if len(data) == 0 {
		responseData = s.SuccessDataWarp(gin.H{})
	} else {
		responseData = s.SuccessDataWarp(data[0])
	}

	s.C.Abort()
	s.C.PureJSON(httpStatusCode, responseData)

	return LogT{
		IsError:        false,
		HTTPStatusCode: httpStatusCode,
		ResponseData:   responseData,
	}
}
