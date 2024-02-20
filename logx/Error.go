package logx

import (
	"os"

	"github.com/rs/zerolog"
)

// 单例
var (
	errorEvent *zerolog.Event
)

// InitWriter_Error
//
// ex:
//
//	logx.InitWriter_Error(filePath, logx.DefaultNewWriterConfig(1, 10, 10))
func InitWriter_Error(filePath string, errorWriterConfig NewWriterConfig) {
	w := NewWriter(filePath+".error.log", errorWriterConfig)
	l := zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	// 存储单例
	errorEvent = l.Error()
}

// Error
//
// ex:
//
//	logx.Error().Msg("")
//	logx.Error().Any("data", data).Send()
func Error() *zerolog.Event {
	if errorEvent != nil {
		return errorEvent
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	return l.Error()
}
