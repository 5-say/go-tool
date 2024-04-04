package logx

import "io"

// 初始化日志，使用默认配置快速初始化
//
// e.g.
//
//	logx.InitDefault("logs/app", "Asia/Shanghai", useAsync, "info", "error", "debug", "gorm", "gin")
//	logx.InitDefault("logs/app", "Asia/Shanghai", useAsync, "info", "error", "debug", "gorm", "nano")
func InitDefault(filePath, locationName string, useAsync bool, tags ...string) {
	loggingLocationName = locationName
	w := func(tag string) io.Writer {
		return DefaultWriter(filePath+"."+tag+".log", useAsync)
	}
	for _, tag := range tags {
		switch tag {
		case "info":
			InitInfoWriter(w(tag))
		case "error":
			InitErrorWriter(w(tag))
		case "debug":
			InitDebugWriter(w(tag))
		case "gin":
			InitGinWriter(w(tag))
		case "gorm":
			InitGormWriter(w(tag))
		case "nano":
			InitNanoWriter(w(tag))
		}
	}
}
