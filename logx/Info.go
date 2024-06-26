package logx

import (
	"io"
	"os"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog"
)

// 单例
var (
	infoLogger *zerolog.Logger
)

// InitInfoWriter
//
// e.g.
//
//	logx.InitInfoWriter(logx.DefaultWriter("info.log", true))
func InitInfoWriter(writer io.Writer) {
	l := zerolog.New(zerolog.ConsoleWriter{Out: writer, NoColor: true, FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName)}).With().Timestamp().Caller().Logger()
	// 存储单例
	infoLogger = &l
}

// Info
//
// e.g.
//
//	logx.Info().Msg("")
//	logx.Info().Any("data", data).Send()
func Info() *zerolog.Event {
	if infoLogger != nil {
		return infoLogger.Info()
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName)}).With().Timestamp().Caller().Logger()
	return l.Info()
}
