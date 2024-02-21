package tool

import (
	"time"

	"github.com/rs/zerolog"
)

// ZerologFormatTimestamp ..
func ZerologFormatTimestamp(loggingLocationName string) zerolog.Formatter {

	return func(i interface{}) string {
		t, err := time.Parse(zerolog.TimeFieldFormat, i.(string))
		if err != nil {
			return err.Error()
		}
		location, err := time.LoadLocation(loggingLocationName)
		if err != nil {
			return err.Error()
		}
		return "\r\n" + t.In(location).Format("MST 2006-01-02 15:04:05")
	}
}
