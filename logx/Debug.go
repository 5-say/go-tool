package logx

import (
	"io"
	"os"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog"
)

// 单例
var (
	debugLogger *zerolog.Logger
)

// InitDebugWriter
//
// ex:
//
//	logx.InitDebugWriter(logx.DefaultWriter("debug.log", true))
func InitDebugWriter(writer io.Writer) {
	l := zerolog.New(zerolog.ConsoleWriter{Out: writer, NoColor: true, FormatTimestamp: tool.ZerologFormatTimestamp(LoggingLocationName)}).With().Timestamp().Caller().Logger()
	// 存储单例
	debugLogger = &l
}

// Debug
//
// ex:
//
//	logx.Debug().Msg("")
//	logx.Debug().Any("data", data).Send()
func Debug() *zerolog.Event {
	if debugLogger != nil {
		return debugLogger.Debug()
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, FormatTimestamp: tool.ZerologFormatTimestamp(LoggingLocationName)}).With().Timestamp().Caller().Logger()
	return l.Debug()
}
