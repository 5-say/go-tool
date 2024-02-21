package tool

import (
	"io"
	"time"

	"github.com/rs/zerolog/diode"
)

// 包装异步写入器
//
//	writer       io.Writer
//	bufferSize   int           // 生产者缓冲区大小
//	pollInterval time.Duration // 消费者轮询间隔
//	alerter      diode.Alerter // 丢弃告警函数
func WrapAsyncWriter(writer io.Writer, bufferSize int, pollInterval time.Duration, alerter diode.Alerter) io.Writer {
	// 线程安全、无锁、无阻塞
	return diode.NewWriter(writer, bufferSize, pollInterval, alerter)
}
