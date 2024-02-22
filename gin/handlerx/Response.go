package handlerx

import "github.com/5-say/go-tool/logx"

// 自定义响应
type Response struct{}

// Success ..
//
// ex:
//
//	return response.Success()
//	return response.Success(400)
//	return response.Success("success message")
func (Response) Success(t ...any) *ResponseBag {
	var r = ResponseBag{httpStatusCode: 200}
	r.data = struct{}{}
	if len(t) == 1 {
		switch v := t[0].(type) {
		case int:
			r.Code(v)
		case string:
			r.Message(v)
		}
	}
	return &r
}

// Error ..
//
// ex:
//
//	return response.Error()
//	return response.Error(500)
//	return response.Error(err)
//	return response.Error("error message")
func (Response) Error(t ...any) *ResponseBag {
	var r = ResponseBag{httpStatusCode: 422}
	r.data = struct{}{}
	if len(t) == 1 {
		switch v := t[0].(type) {
		case int:
			r.Code(v)
		case error:
			logx.Error().Err(v).Send()
		case string:
			r.Message(v)
		}
	}
	return &r
}
