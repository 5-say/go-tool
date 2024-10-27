package responsex

// Success ..
//
// e.g.
//
//	r.Success()
//	r.Success(data any)
//	r.Success(message string)
//	r.Success(message string, data any)
//	r.Success(messageFormat []any)
//	r.Success(messageFormat []any, data any)
func (s ResponseT) Success(data ...any) LogT {
	return s.SuccessCode(s.DefaultSuccessHTTPCode, data...)
}

// SuccessCode ..
//
// e.g.
//
//	r.SuccessCode(httpStatusCode int)
//	r.SuccessCode(httpStatusCode int, data any)
//	r.SuccessCode(httpStatusCode int, message string)
//	r.SuccessCode(httpStatusCode int, message string, data any)
//	r.SuccessCode(httpStatusCode int, messageFormat []any)
//	r.SuccessCode(httpStatusCode int, messageFormat []any, data any)
func (s ResponseT) SuccessCode(httpStatusCode int, data ...any) LogT {
	// 构造响应数据
	var responseData any
	if len(data) == 0 {
		responseData = s.SuccessDataWarp([]any{"success"}, nil)
	} else if len(data) == 1 {
		switch v := data[0].(type) {
		case string:
			responseData = s.SuccessDataWarp([]any{v}, nil)
		case []any:
			responseData = s.SuccessDataWarp(v, nil)
		default:
			responseData = s.SuccessDataWarp([]any{"success"}, v)
		}
	} else if len(data) == 2 {
		switch v := data[0].(type) {
		case string:
			responseData = s.SuccessDataWarp([]any{v}, data[1])
		case []any:
			responseData = s.SuccessDataWarp(v, data[1])
		}
	}

	s.C.Abort()
	s.C.PureJSON(httpStatusCode, responseData)

	return LogT{
		IsError:        false,
		HTTPStatusCode: httpStatusCode,
		ResponseData:   responseData,
	}
}
