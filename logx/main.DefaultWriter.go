package logx

import "io"

// DefaultWriter ..
//
// e.g.
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
