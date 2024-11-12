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
//	statusCode, err := httpx.PostJSON(urlStr, jsonParams, headers).BindJSON(&apiResult)
func (s ResponseT) BindJSON(obj any) (statusCode int, err error) {
	if s.Error != nil {
		return s.StatusCode, s.Error
	}
	return s.StatusCode, json.Unmarshal(s.Body, obj)
}
