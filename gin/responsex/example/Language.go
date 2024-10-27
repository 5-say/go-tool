package example

import "github.com/5-say/go-tool/gin/responsex"

var Language = responsex.LanguageT{
	Direct: func() bool {
		return false
	},
	Use: "pt",
	Map: map[string]map[string]string{
		"pt": {
			"最小提现金额：%v": "Valor mínimo de retirada: %v",
		},
	},
}
