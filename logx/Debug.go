package logx

import (
	"time"

	"github.com/rs/zerolog"
)

// Debug
//
// ex:
//
//	logx.Debug().Msg("")
//	logx.Debug().Any("data", data).Send()
func Debug() *zerolog.Event {
	timeStr := time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	return DebugLogger.Log().Str("time", timeStr).Caller(1)
}
