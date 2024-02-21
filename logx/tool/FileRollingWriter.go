package tool

import (
	"io"

	"github.com/natefinch/lumberjack"
)

// 文件滚动写入器
//
//	fileName     string // 包含完整路径的文件名
//	MaxMegabytes int    // 单文件最大存储容量（兆字节）
//	MaxDays      int    // 最大存储天数
//	MaxBackups   int    // 最大备份文件数量
//	Compress     bool   // 备份文件是否压缩
func FileRollingWriter(fileName string, maxMegabytes, maxDays, maxBackups int, compress bool) io.Writer {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxMegabytes,
		MaxAge:     maxDays,
		MaxBackups: maxBackups,
		Compress:   compress,
	}
}
