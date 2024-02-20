package logx

import (
	"os"

	"github.com/rs/zerolog"
)

// 单例
var (
	infoLogger *zerolog.Logger
)

// InitWriter_Info
//
// ex:
//
//	logx.InitWriter_Info(filePath, logx.DefaultNewWriterConfig(1, 10, 10))
func InitWriter_Info(filePath string, infoWriterConfig NewWriterConfig) {
	w := NewWriter(filePath+".info.log", infoWriterConfig)
	l := zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	// 存储单例
	infoLogger = &l
}

// Info
//
// ex:
//
//	logx.Info().Msg("")
//	logx.Info().Any("data", data).Send()
func Info() *zerolog.Event {
	if infoLogger != nil {
		return infoLogger.Info()
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	return l.Info()
}
