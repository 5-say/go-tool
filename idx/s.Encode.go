package idx

import (
	"math/rand"
	"strconv"
	"time"
)

// 编码
func (s TootT) Encode(id uint64) string {
	// 初始化随机数生成器
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 原始字符串
	original := strconv.FormatUint(id+s.Offset, s.Base)

	// 需要插入的字符
	insertChars := s.FillChars

	// 原始长度
	originalLength := len(original)

	// 目标长度
	targetLength := s.MinLenth

	// 填充长度
	fillLength := targetLength - originalLength
	if fillLength < 0 {
		fillLength = 0
	}

	// 创建一个切片来存储结果字符串的 rune
	result := make([]rune, 0, targetLength)

	// 将原始字符串的 rune 添加到结果切片中
	for _, r := range original {
		result = append(result, r)

		// 随机插入填充字符
		if fillLength > 0 {
			// 随机选择一个插入位置（包括当前位置之后）
			insertPos := rand.Intn(len(result) + 1)

			// 随机选择一个要插入的字符
			insertChar := insertChars[rand.Intn(len(insertChars))]

			// 在指定位置插入字符
			result = append(result[:insertPos], append([]rune{insertChar}, result[insertPos:]...)...)

			// 标记
			fillLength--
		}
	}

	// 如果结果切片长度仍然不足目标长度，则在末尾补足随机字符
	for len(result) < targetLength {
		// 随机选择一个要插入的字符
		insertChar := insertChars[rand.Intn(len(insertChars))]
		result = append(result, insertChar)
	}

	// 创建一个切片来存储混淆后的字符串的 rune
	mixedResult := make([]rune, 0, targetLength)

	// 混淆结果切片
	for _, r := range result {
		mixedR, ok := s.MixMap[r]
		if !ok {
			mixedR = r
		}
		mixedResult = append(mixedResult, mixedR)
	}

	// 将结果切片转换回字符串
	return string(mixedResult) // 允许超出长度
}
