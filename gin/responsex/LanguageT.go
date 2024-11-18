package responsex

import (
	"fmt"

	"git.samba7.com/hb-game-lobby/gl-dev-db/platform/dao/modelPlatform"
	"github.com/5-say/go-tool/logx"
	"gorm.io/gorm"
)

type LanguageT struct {
	Direct func() bool                  // 是否直接输出原始消息
	Use    string                       // 当前使用的语言名称
	Map    map[string]map[string]string // 语言包定义
	DB     func() *gorm.DB              // 支持数据库获取语音包
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

func (s LanguageT) GetFromDB(messageFormat []any) (message string) {
	var (
		db       = s.DB()
		format   = messageFormat[0].(string)
		language map[string]any
	)

	db.Model(&modelPlatform.Language{}).Where("format = ?", format).Find(&language)
	if len(language) == 0 {
		if err := db.Model(&modelPlatform.Language{}).Create(map[string]any{
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
