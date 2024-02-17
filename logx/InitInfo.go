package logx

import (
	"time"

	"github.com/rs/zerolog/diode"
)

// InitInfo
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
//	logx.InitInfo("demo", 3, 30, 100, false, 1000, 0, nil)
func InitInfo(fileName string, maxMegabytes, maxDays, maxBackups int, compress bool, diodeBufferSize int, diodePollInterval time.Duration, diodeAlerter diode.Alerter) {
	InfoLogger = InitLogger(fileName+".info.log", maxMegabytes, maxDays, maxBackups, compress, diodeBufferSize, diodePollInterval, diodeAlerter)
}
