package logx

import (
	"os"

	"github.com/rs/zerolog"
)

// 单例
var (
	debugLogger *zerolog.Logger
)

// InitWriter_Debug
//
// ex:
//
//	logx.InitWriter_Debug(filePath, logx.DefaultNewWriterConfig(1, 10, 10))
func InitWriter_Debug(filePath string, debugWriterConfig NewWriterConfig) {
	w := NewWriter(filePath+".debug.log", debugWriterConfig)
	l := zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
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
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	return l.Debug()
}
