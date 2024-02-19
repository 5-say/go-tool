package logx

import (
	"os"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

// 单例
var (
	gormWriter logger.Writer
)

func InitWriter_Gorm(filePath string, gormWriterConfig NewWriterConfig) {
	w := NewWriter(filePath+".gorm.log", gormWriterConfig)
	l := zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: true, FormatTimestamp: FormatTimestamp}).With().Timestamp().Logger()
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
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, FormatTimestamp: FormatTimestamp}).With().Timestamp().Logger()
	config.Colorful = true
	return logger.New(&l, config)
}
