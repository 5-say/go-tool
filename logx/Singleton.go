package logx

import (
	"github.com/rs/zerolog"
)

// 单例
var (
	InfoLogger  zerolog.Logger
	ErrorLogger zerolog.Logger
	DebugLogger zerolog.Logger
)
