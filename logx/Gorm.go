package logx

import (
	"io"
	"os"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

// 单例
var (
	gormWriter logger.Writer
)

// InitGormWriter
//
// ex:
//
//	logx.InitGormWriter(logx.DefaultWriter("gorm.log", true))
func InitGormWriter(writer io.Writer) {
	l := zerolog.New(zerolog.ConsoleWriter{Out: writer, NoColor: true, FormatTimestamp: tool.ZerologFormatTimestamp(LoggingLocationName)}).With().Timestamp().Logger()
	// 存储单例
	gormWriter = &l
}

// GormLogger
//
// ex:
//
//	logx.GormLogger(logger.Config{
//		SlowThreshold:             200 * time.Millisecond,
//		IgnoreRecordNotFoundError: false,
//		ParameterizedQueries:      false,
//		LogLevel:                  logger.Warn,
//	})
func GormLogger(config logger.Config) logger.Interface {
	if gormWriter != nil {
		config.Colorful = false
		return logger.New(gormWriter, config)
	}
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: tool.ZerologFormatTimestamp(LoggingLocationName)}).With().Timestamp().Logger()
	config.Colorful = true
	return logger.New(&l, config)
}
