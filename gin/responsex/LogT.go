package responsex

import "github.com/5-say/go-tool/logx"

type LogT struct {
	IsError        bool
	HTTPStatusCode int
	ResponseData   interface{}
}

// Log ..
//
// e.g.
//
//	r.Error().Log()
func (s LogT) Log() LogT {
	if s.IsError {
		logx.Error().CallerSkipFrame(1).Int("HTTPStatusCode", s.HTTPStatusCode).Any("ResponseData", s.ResponseData).Send()
	} else {
		logx.Info().CallerSkipFrame(1).Int("HTTPStatusCode", s.HTTPStatusCode).Any("ResponseData", s.ResponseData).Send()
	}
	return s
}

// Logf ..
//
// e.g.
//
//	r.Error().Logf(format string, v ...interface{})
func (s LogT) Logf(format string, v ...interface{}) LogT {
	if s.IsError {
		logx.Error().CallerSkipFrame(1).Msgf(format, v...)
	} else {
		logx.Info().CallerSkipFrame(1).Msgf(format, v...)
	}
	return s
}
