package handlerx

import "github.com/5-say/go-tool/logx"

// 单例
var Language = LanguageType{
	Debug: false,
	Force: false,
	Use:   "",
	Map:   make(map[string]map[string]string),
}

type LanguageType struct {
	Debug bool                         // 调试模式，直接输出原始消息
	Force bool                         // 是否强制启用语言包
	Use   string                       // 当前使用的语言名称
	Map   map[string]map[string]string // 语言包定义
}

func (s LanguageType) Set(languageName, key, value string) {
	if _, ok := s.Map[languageName]; !ok {
		s.Map[languageName] = make(map[string]string)
	}
	s.Map[languageName][key] = value
}

func (s LanguageType) Get(key string) (value string) {
	if s.Debug {
		return key
	}
	value, ok := s.Map[s.Use][key]
	if !ok {
		if s.Force {
			value = "..."
		} else {
			value = key
		}
		logx.Debug().Msg("缺少语言文件 languageName: " + s.Use + " key: " + key)
	}
	return
}
