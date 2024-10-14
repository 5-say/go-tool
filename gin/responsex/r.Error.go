package responsex

import "github.com/gin-gonic/gin"

// Error ..
//
// e.g.
//
//	r.Error()
//	r.Error(data interface{})
func (s ResponseT) Error(data ...interface{}) LogT {
	return s.ErrorCode(s.DefaultErrorHTTPCode, data...)
}

// ErrorCode ..
//
// e.g.
//
//	r.ErrorCode(httpStatusCode int)
//	r.ErrorCode(httpStatusCode int, data interface{})
func (s ResponseT) ErrorCode(httpStatusCode int, data ...interface{}) LogT {
	// 构造响应数据
	var responseData interface{}
	if len(data) == 0 {
		responseData = s.ErrorDataWarp(gin.H{})
	} else {
		responseData = s.ErrorDataWarp(data[0])
	}

	s.C.Abort()
	s.C.PureJSON(httpStatusCode, responseData)

	return LogT{
		IsError:        true,
		HTTPStatusCode: httpStatusCode,
		ResponseData:   responseData,
	}
}
