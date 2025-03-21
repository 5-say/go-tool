// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package modelPlatform

const TableNameLanguage = "language"

// Language 语言定义表
type Language struct {
	ID     uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`        // ID
	Format string `gorm:"column:format;type:varchar(255);uniqueIndex:format,priority:1;comment:原始格式" json:"format"` // 原始格式
	En     string `gorm:"column:en;type:varchar(255);default:...;comment:英语" json:"en"`                             // 英语
	Pt     string `gorm:"column:pt;type:varchar(255);default:...;comment:葡萄牙语" json:"pt"`                           // 葡萄牙语
	Es     string `gorm:"column:es;type:varchar(255);default:...;comment:西班牙语" json:"es"`                           // 西班牙语
}

// TableName Language's table name
func (*Language) TableName() string {
	return TableNameLanguage
}
