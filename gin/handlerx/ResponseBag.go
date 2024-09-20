package handlerx

import "github.com/5-say/go-tool/logx"

// ResponseBag ..
type ResponseBag struct {
	httpStatusCode int    // 当前响应的 http 状态码
	mark           string // 当前响应的类型标记（可以是任意字符串，提供给前端区分处理不同的响应类型）
	message        string // 可以直接展示给用户的提示信息
	resource       any    // 当前响应的主体资源
}

// HTTPCode ..
func (r *ResponseBag) HTTPCode(httpStatusCode int) *ResponseBag {
	r.httpStatusCode = httpStatusCode
	return r
}

// Mark ..
func (r *ResponseBag) Mark(mark string) *ResponseBag {
	r.mark = mark
	return r
}

// Resource ..
func (r *ResponseBag) Resource(resource any) *ResponseBag {
	r.resource = resource
	return r
}

// Message ..
func (r *ResponseBag) Message(message string) *ResponseBag {
	r.message = message
	return r
}

// Log ..
func (r *ResponseBag) Log(t any) *ResponseBag {
	switch v := t.(type) {
	case error:
		logx.Debug().CallerSkipFrame(1).Err(v).Send()
	default:
		logx.Debug().CallerSkipFrame(1).Any("", v).Send()
	}
	return r
}
