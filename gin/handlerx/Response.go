package handlerx

// 自定义响应
type Response struct{}

// Success ..
//
// e.g.
//
//	return r.Success()
//	return r.Success(400)
//	return r.Success("success message")
//	return r.Success().Resource(gin.H{})
func (Response) Success(t ...any) *ResponseBag {
	var r = ResponseBag{httpStatusCode: 200, mark: "success"}
	r.resource = struct{}{}
	if len(t) == 1 {
		switch v := t[0].(type) {
		case int:
			r.HTTPCode(v)
		case string:
			r.Message(v)
		}
	}
	return &r
}

// Error ..
//
// e.g.
//
//	return r.Error()
//	return r.Error(500)
//	return r.Error("error message")
//	return r.Error().Log(err)
func (Response) Error(t ...any) *ResponseBag {
	var r = ResponseBag{httpStatusCode: 422, mark: "error"}
	r.resource = struct{}{}
	if len(t) == 1 {
		switch v := t[0].(type) {
		case int:
			r.HTTPCode(v)
		case string:
			r.Message(v)
		}
	}
	return &r
}
