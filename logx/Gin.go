package logx

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// 单例
var (
	ginWriter io.Writer
)

// InitWriter_Gin
//
// ex:
//
//	logx.InitWriter_Gin(filePath, logx.DefaultNewWriterConfig(1, 10, 10))
func InitWriter_Gin(filePath string, ginWriterConfig NewWriterConfig) {
	w := NewWriter(filePath+".gin.log", ginWriterConfig)
	l := zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
	// 存储单例
	ginWriter = l
}

// GinWriter
//
// ex:
//
//	if !isTestMode {
//		gin.DefaultWriter = logx.GinWriter()
//		gin.DefaultErrorWriter = logx.GinWriter()
//	}
//	r := gin.Default()
func GinWriter() io.Writer {
	if ginWriter != nil {
		return ginWriter
	}
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: FormatTimestamp}).With().Timestamp().Caller().Logger()
}
