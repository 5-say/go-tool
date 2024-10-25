package responsex

// Error ..
//
// e.g.
//
//	r.Error()
//	r.Error(data any)
//	r.Error(err error)
//	r.Error(err error, data any)
//	r.Error(message string)
//	r.Error(message string, data any)
func (s ResponseT) Error(data ...interface{}) LogT {
	return s.ErrorCode(s.DefaultErrorHTTPCode, data...)
}

// ErrorCode ..
//
// e.g.
//
//	r.ErrorCode(httpStatusCode int, )
//	r.ErrorCode(httpStatusCode int, data any)
//	r.ErrorCode(httpStatusCode int, err error)
//	r.ErrorCode(httpStatusCode int, err error, data any)
//	r.ErrorCode(httpStatusCode int, message string)
//	r.ErrorCode(httpStatusCode int, message string, data any)
func (s ResponseT) ErrorCode(httpStatusCode int, data ...interface{}) LogT {
	// 构造响应数据
	var responseData any
	if len(data) == 0 {
		responseData = s.ErrorDataWarp("error", nil, nil)
	} else if len(data) == 1 {
		switch v := data[0].(type) {
		case error:
			responseData = s.ErrorDataWarp("error", nil, v)
		case string:
			responseData = s.ErrorDataWarp(v, nil, nil)
		default:
			responseData = s.ErrorDataWarp("error", v, nil)
		}
	} else if len(data) == 2 {
		switch v := data[0].(type) {
		case error:
			responseData = s.ErrorDataWarp("error", data[1], v)
		case string:
			responseData = s.ErrorDataWarp(v, data[1], nil)
		}
	}

	s.C.Abort()
	s.C.PureJSON(httpStatusCode, responseData)

	return LogT{
		IsError:        true,
		HTTPStatusCode: httpStatusCode,
		ResponseData:   responseData,
	}
}
