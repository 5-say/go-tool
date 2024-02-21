package logx

import "io"

// LoggingLocationName
//
// ex:
//
//	logx.LoggingLocationName = "Asia/Shanghai"
var LoggingLocationName = "UTC"

// DefaultWriter ..
//
// ex:
//
//	logx.DefaultWriter("logs/demo.log", true)
func DefaultWriter(fileName string, useAsync bool) io.Writer {
	config := WriterConfig{
		Rolling: WriterRollingConfig{
			MaxMegabytes: 1,
			MaxDays:      10,
			MaxBackups:   10,
			Compress:     false,
		},
	}
	if useAsync {
		config.Async = &WriterAsyncConfig{
			BufferSize:   1000,
			PollInterval: 0,
			Alerter:      nil,
		}
	}
	return NewWriter(fileName, config)
}

// 初始化日志
//
// ex:
//
//	logx.InitDefault("logs/app", "Asia/Shanghai", isTestMode, "info", "error", "debug", "gin", "gorm")
func InitDefault(filePath, loggingLocationName string, isTestMode bool, tags ...string) {
	LoggingLocationName = loggingLocationName
	w := func(tag string) io.Writer {
		return DefaultWriter(filePath+"."+tag+".log", !isTestMode)
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
		}
	}
}
