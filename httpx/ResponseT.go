package httpx

import (
	"encoding/json"
)

// ResponseT ..
type ResponseT struct {
	Error      error
	Body       []byte
	StatusCode int
}

// 解析响应，绑定到对象
//
// e.g.
//
//	var obj struct {
//		Status  string
//		Message string
//		Data    map[string]interface{}
//	}
//	_, err := httpx.PostJSON(url, jsonParams, nil).BindJSON(&obj)
func (s ResponseT) BindJSON(obj any) (statusCode int, err error) {
	if s.Error != nil {
		return s.StatusCode, s.Error
	}
	return s.StatusCode, json.Unmarshal(s.Body, obj)
}
