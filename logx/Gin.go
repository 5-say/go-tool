package logx

import (
	"io"
	"os"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog"
)

// 单例
var (
	ginWriter io.Writer
)

// InitGinWriter
//
// e.g.
//
//	logx.InitGinWriter(logx.DefaultWriter("gin.log", true))
func InitGinWriter(writer io.Writer) {
	l := zerolog.New(zerolog.ConsoleWriter{Out: writer, NoColor: true, FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName)}).With().Timestamp().Caller().Logger()
	// 存储单例
	ginWriter = l
}

// GinWriter
//
// e.g.
//
//	gin.DefaultWriter = logx.GinWriter()
//	gin.DefaultErrorWriter = logx.GinWriter()
//	r := gin.Default()
func GinWriter() io.Writer {
	if ginWriter != nil {
		return ginWriter
	}
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: tool.ZerologFormatTimestamp(loggingLocationName)}).With().Timestamp().Caller().Logger()
}
