package logx

import (
	"os"

	"github.com/rs/zerolog"
)

var (
	infoEvent *zerolog.Event
)

func InitWriter_Info(filePath string, infoWriterConfig NewWriterConfig) {
	w := NewWriter(filePath+".info.log", infoWriterConfig)
	l := zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	// 存储单例
	infoEvent = l.Info()
}

// Info
//
// ex:
//
//	logx.Info().Msg("")
//	logx.Info().Any("data", data).Send()
func Info() *zerolog.Event {
	if infoEvent != nil {
		return infoEvent
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	return l.Info()
}
