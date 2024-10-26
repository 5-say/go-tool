package idx

import (
	"strconv"
	"strings"
)

// 解码
func (s TootT) Decode(encryptedStr string) (id uint64, err error) {
	// 翻转混淆字典
	var recoverMap = map[rune]rune{}
	for k, v := range s.MixMap {
		recoverMap[v] = k
	}

	// 恢复混淆字符串
	recoverResult := make([]rune, 0, s.MinLenth)
	for _, r := range encryptedStr {
		recoverR, ok := recoverMap[r]
		if !ok {
			recoverR = r
		}
		recoverResult = append(recoverResult, recoverR)
	}

	// 移除填充字符串
	unfilledResult := string(recoverResult)
	for _, r := range s.FillChars {
		unfilledResult = strings.ReplaceAll(unfilledResult, string(r), "")
	}

	// 进制转换
	decimalInt, err := strconv.ParseInt(unfilledResult, s.Base, 0)
	if err != nil {
		return
	}

	return uint64(decimalInt) - s.Offset, nil
}
