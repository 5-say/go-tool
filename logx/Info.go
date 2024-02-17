package logx

import (
	"time"

	"github.com/rs/zerolog"
)

// Info
//
// ex:
//
//	logx.Info().Msg("")
//	logx.Info().Any("data", data).Send()
func Info() *zerolog.Event {
	timeStr := time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	return InfoLogger.Log().Str("time", timeStr).Caller(1)
}
