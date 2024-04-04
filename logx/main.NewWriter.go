package logx

import (
	"io"
	"time"

	"github.com/5-say/go-tool/logx/tool"
	"github.com/rs/zerolog/diode"
)

// WriterConfig ..
type WriterConfig struct {
	Rolling WriterRollingConfig
	Async   *WriterAsyncConfig
}

// 滚动写入配置
type WriterRollingConfig struct {
	MaxMegabytes int  // 单文件最大存储容量（兆字节）
	MaxDays      int  // 最大存储天数
	MaxBackups   int  // 最大备份文件数量
	Compress     bool // 备份文件是否压缩
}

// 异步写入配置
type WriterAsyncConfig struct {
	BufferSize   int           // 生产者缓冲区大小
	PollInterval time.Duration // 消费者轮询间隔
	Alerter      diode.Alerter // 丢弃告警函数
}

// 实例化日志写入器，支持 滚动 与 异步
func NewWriter(fileName string, c WriterConfig) io.Writer {
	var w = tool.FileRollingWriter(fileName, c.Rolling.MaxBackups, c.Rolling.MaxDays, c.Rolling.MaxMegabytes, c.Rolling.Compress)
	if c.Async != nil {
		return tool.WrapAsyncWriter(w, c.Async.BufferSize, c.Async.PollInterval, c.Async.Alerter)
	}
	return w
}
