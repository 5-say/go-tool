package logx

import (
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

// InitLogger
//
//	fileName          string        日志文件名
//	maxMegabytes      int           单文件最大容量（兆字节）
//	maxDays           int           最大存储天数
//	maxBackups        int           最大备份文件数量
//	compress          bool          备份文件是否压缩
//	diodeBufferSize   int           diode 缓冲区大小
//	diodePollInterval time.Duration diode 轮询时间间隔
//	diodeAlerter      diode.Alerter diode 告警函数
//
// ex:
//
//	logger := logx.InitLogger("demo", 3, 30, 100, false, 1000, 0, nil)
func InitLogger(fileName string, maxMegabytes, maxDays, maxBackups int, compress bool, diodeBufferSize int, diodePollInterval time.Duration, diodeAlerter diode.Alerter) zerolog.Logger {

	// 日志拆分
	rollingWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxMegabytes,
		MaxAge:     maxDays,
		MaxBackups: maxBackups,
		Compress:   compress,
	}

	// 线程安全、无锁、无阻塞
	diodeWriter := diode.NewWriter(rollingWriter, diodeBufferSize, diodePollInterval, diodeAlerter)

	// 单例模式
	return zerolog.New(diodeWriter)
}
