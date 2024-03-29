package handlerx

import "github.com/5-say/go-tool/logx"

// 自定义响应
type Response struct{}

// Success ..
//
// e.g.
//
//	return response.Success()
//	return response.Success(400)
//	return response.Success("success message")
func (Response) Success(t ...any) *ResponseBag {
	var r = ResponseBag{httpStatusCode: 200, status: "success"}
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
// e.g.
//
//	return response.Error()
//	return response.Error(500)
//	return response.Error(err) // 接收 error 类型，只打印日志，不输出到 message
//	return response.Error("error message")
func (Response) Error(t ...any) *ResponseBag {
	var r = ResponseBag{httpStatusCode: 422, status: "error"}
	r.data = struct{}{}
	if len(t) == 1 {
		switch v := t[0].(type) {
		case int:
			r.Code(v)
		case error:
			logx.Error().CallerSkipFrame(1).Err(v).Send()
		case string:
			r.Message(v)
		}
	}
	return &r
}
