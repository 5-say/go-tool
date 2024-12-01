package example

import (
	"github.com/5-say/go-tool/gin/responsex"
)

var Language = responsex.LanguageT{
	Direct: func() bool {
		return false
	},
	Use: "pt",
	Map: map[string]map[string]string{
		"pt": {
			"":          "%s",
			"success":   "Success",
			"error":     "Error",
			"参数错误":      "Erro de parâmetro",
			"操作失败":      "Falha na operação",
			"最小提现金额：%v": "Valor mínimo de retirada: %v",
		},
	},
	// Model: func() *gorm.DB {
	// 	return DB().Model(&model.Language{})
	// },
}
