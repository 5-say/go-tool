package responsex

// Error ..
//
// e.g.
//
//	r.Error()
//	r.Error(message string)
func (s ResponseT) Error(message ...string) LogT {
	return s.ErrorCode(s.DefaultErrorHTTPCode, message...)
}

// ErrorCode ..
//
// e.g.
//
//	r.ErrorCode(httpStatusCode int)
//	r.ErrorCode(httpStatusCode int, message string)
func (s ResponseT) ErrorCode(httpStatusCode int, message ...string) LogT {
	// 构造响应数据
	var responseData interface{}
	if len(message) == 0 {
		responseData = s.ErrorDataWarp("")
	} else {
		responseData = s.ErrorDataWarp(message[0])
	}

	s.C.Abort()
	s.C.PureJSON(httpStatusCode, responseData)

	return LogT{
		IsError:        true,
		HTTPStatusCode: httpStatusCode,
		ResponseData:   responseData,
	}
}
