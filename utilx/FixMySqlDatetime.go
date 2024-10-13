package utilx

import (
	"time"

	"github.com/5-say/go-tool/change"
)

// FixMySqlDatetime ..
// 规避 MySQL 数据库 DATETIME 字段类型存储时，毫秒位四舍五入的问题
func FixMySqlDatetime(timeObj time.Time) (fixedTimestamp int64) {
	originalTimestamp := timeObj.Unix()
	if change.ForceTo(timeObj.Format(".000")[1:], 0) >= 500 {
		return originalTimestamp + 1
	}
	return originalTimestamp
}
