package handlerx

import "github.com/5-say/go-tool/logx"

// ResponseBag ..
type ResponseBag struct {
	httpStatusCode int
	data           any
	mark           string // 提供给前端程序内部使用的标记位
	status         string // 当前响应状态：success、error（初始化后不可改变）
	message        string // 可以直接展示给用户的信息
}

// Code ..
func (r *ResponseBag) Code(httpStatusCode int) *ResponseBag {
	r.httpStatusCode = httpStatusCode
	return r
}

// Mark ..
func (r *ResponseBag) Mark(mark string) *ResponseBag {
	r.mark = mark
	return r
}

// Data ..
func (r *ResponseBag) Data(data any) *ResponseBag {
	r.data = data
	return r
}

// Message ..
func (r *ResponseBag) Message(message string) *ResponseBag {
	r.message = message
	return r
}

// Debug ..
func (r *ResponseBag) Debug(i interface{}) *ResponseBag {
	logx.Debug().CallerSkipFrame(1).Any("", i).Send()
	return r
}
