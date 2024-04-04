package logx

import (
	"io"
	"os"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog"
)

// 单例
var (
	errorLogger *zerolog.Logger
)

// InitErrorWriter
//
// e.g.
//
//	logx.InitErrorWriter(logx.DefaultWriter("error.log", true))
func InitErrorWriter(writer io.Writer) {
	l := zerolog.New(zerolog.ConsoleWriter{Out: writer, NoColor: true, FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName)}).With().Timestamp().Caller().Logger()
	// 存储单例
	errorLogger = &l
}

// Error
//
// e.g.
//
//	logx.Error().Msg("")
//	logx.Error().Any("data", data).Send()
func Error() *zerolog.Event {
	if errorLogger != nil {
		return errorLogger.Error()
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName)}).With().Timestamp().Caller().Logger()
	return l.Error()
}
