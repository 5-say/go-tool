package logx

import (
	"time"

	"github.com/rs/zerolog"
)

// Error
//
// ex:
//
//	logx.Error().Msg("")
//	logx.Error().Any("data", data).Send()
func Error() *zerolog.Event {
	timeStr := time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	return ErrorLogger.Log().Str("time", timeStr).Caller(1)
}
