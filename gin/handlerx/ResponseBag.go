package handlerx

import "github.com/5-say/go-tool/logx"

// ResponseBag ..
type ResponseBag struct {
	httpStatusCode int
	mark           string
	data           any
	message        string
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
