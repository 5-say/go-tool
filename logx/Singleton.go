package logx

import (
	"time"

	"github.com/rs/zerolog"
)

// 日志记录时区
//
// ex:
//
//	logx.LoggingLocationName = "Asia/Shanghai"
var LoggingLocationName = "UTC"

func FormatTimestamp(i interface{}) string {
	t, err := time.Parse(zerolog.TimeFieldFormat, i.(string))
	if err != nil {
		return err.Error()
	}
	location, err := time.LoadLocation(LoggingLocationName)
	if err != nil {
		return err.Error()
	}
	return "\r\n" + t.In(location).Format("MST 2006-01-02 15:04:05")
}
