package responsex

import (
	"fmt"

	"github.com/5-say/go-tool/logx"
)

type LanguageT struct {
	Direct func() bool                  // 是否直接输出原始消息
	Use    string                       // 当前使用的语言名称
	Map    map[string]map[string]string // 语言包定义
}

func (s LanguageT) Get(messageFormat []any) (message string) {
	var format = messageFormat[0].(string)
	if s.Direct() {
		return fmt.Sprintf(format, messageFormat[1:]...)
	}
	message, ok := s.Map[s.Use][format]
	if !ok {
		logx.Debug().CallerSkipFrame(4).Msgf(`缺少语言文件 languageName:"%s" key:"%s"`, s.Use, format)
		return "..."
	}
	return fmt.Sprintf(message, messageFormat[1:]...)
}
