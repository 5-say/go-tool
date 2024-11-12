package httpx

import (
	"encoding/json"
)

// 解析响应，绑定到对象
//
//	obj any
//
//	return
//
//	statusCode int
//	err        error
//
// e.g.
//
//	statusCode, body, err := httpx.PostJSON(urlStr, jsonParams, headers).BindJSON(&apiResult)
func (s ResponseT) BindJSON(obj any) (statusCode int, body string, err error) {
	if s.Error != nil {
		return s.StatusCode, string(s.Body), s.Error
	}
	return s.StatusCode, string(s.Body), json.Unmarshal(s.Body, obj)
}
