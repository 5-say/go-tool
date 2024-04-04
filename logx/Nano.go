package logx

import (
	"fmt"
	"io"
	"os"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog"
)

type nanoLogger struct{}

func (nanoLogger) Println(v ...interface{}) {
	if e := nanoWriter.Info(); e.Enabled() {
		e.CallerSkipFrame(1).Msg(fmt.Sprint(v...))
	}
}

func (nanoLogger) Fatal(v ...interface{}) {
	nanoWriter.Fatal().Msg(fmt.Sprint(v...))
}

func (nanoLogger) Fatalf(format string, v ...interface{}) {
	nanoWriter.Fatal().Msgf(format, v...)
}

// 单例
var (
	nanoWriter *zerolog.Logger
)

// InitNanoWriter
//
// e.g.
//
//	logx.InitNanoWriter(logx.DefaultWriter("nano.log", true))
func InitNanoWriter(writer io.Writer) {
	l := zerolog.New(zerolog.ConsoleWriter{
		Out: writer, NoColor: true,
		FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName),
	}).With().Timestamp().Caller().CallerWithSkipFrameCount(3).Logger()
	// 存储单例
	nanoWriter = &l
}

// NanoLogger
//
// e.g.
//
//	nano.WithLogger(logx.NanoLogger())
func NanoLogger() nanoLogger {
	if nanoWriter == nil {
		l := zerolog.New(zerolog.ConsoleWriter{
			Out: os.Stdout, NoColor: false,
			FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName),
		}).With().Timestamp().Caller().CallerWithSkipFrameCount(3).Logger()
		// 存储单例
		nanoWriter = &l
	}
	return nanoLogger{}
}
