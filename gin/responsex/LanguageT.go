package responsex

import (
	"fmt"

	"github.com/5-say/go-tool/logx"
	"gorm.io/gorm"
)

type LanguageT struct {
	Direct func() bool                  // 是否直接输出原始消息
	Use    string                       // 当前使用的语言名称
	Map    map[string]map[string]string // 基础语言包定义
	Model  func() *gorm.DB              // 支持数据库获取语言包
}

// 直接使用基础语言包
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

// 从数据库获取
func (s LanguageT) GetFromDB(messageFormat []any) (message string) {
	var (
		format   = messageFormat[0].(string)
		language map[string]any
	)

	s.Model().Where("format = ?", format).Find(&language)
	if len(language) == 0 {
		if err := s.Model().Create(map[string]any{
			"format": format,
		}).Error; err != nil {
			logx.Debug().CallerSkipFrame(4).Err(err)
		}

		if s.Direct() {
			return fmt.Sprintf(format, messageFormat[1:]...)
		}

		return "..."
	}

	return fmt.Sprintf(language[s.Use].(string), messageFormat[1:]...)
}
