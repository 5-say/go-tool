package logx

import (
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog/diode"
)

// NewWriter
//
//	fileName string               日志文件名
//	config   logx.NewWriterConfig 文件日志配置
//
// ex:
//
//	logx.NewWriter("demo.log", logx.)
func NewWriter(fileName string, config NewWriterConfig) diode.Writer {
	var c = config

	// 日志拆分
	rollingWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    c.MaxMegabytes,
		MaxAge:     c.MaxDays,
		MaxBackups: c.MaxBackups,
		Compress:   c.Compress,
	}

	// 线程安全、无锁、无阻塞
	return diode.NewWriter(rollingWriter, c.DiodeBufferSize, c.DiodePollInterval, c.DiodeAlerter)
}

// NewWriterConfig ..
type NewWriterConfig struct {
	MaxMegabytes      int           // 单文件最大容量（兆字节）
	MaxDays           int           // 最大存储天数
	MaxBackups        int           // 最大备份文件数量
	Compress          bool          // 备份文件是否压缩
	DiodeBufferSize   int           // diode 缓冲区大小
	DiodePollInterval time.Duration // diode 轮询时间间隔
	DiodeAlerter      diode.Alerter // diode 告警函数
}

// DefaultNewWriterConfig ..
func DefaultNewWriterConfig(maxMegabytes, maxDays, maxBackups int) NewWriterConfig {
	return NewWriterConfig{
		MaxMegabytes:      maxMegabytes,
		MaxDays:           maxDays,
		MaxBackups:        maxBackups,
		Compress:          false,
		DiodeBufferSize:   1000,
		DiodePollInterval: 0,
		DiodeAlerter:      nil,
	}
}
