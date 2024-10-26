package idx

type TootT struct {
	MinLenth  int           // 生成字符串的最小长度
	Base      int           // 进制基数，2 <= base <= 36
	Offset    uint64        // 初始偏移量
	FillChars []rune        // 填充字符，需超出进制覆盖
	MixMap    map[rune]rune // 混淆字典，定义时确保闭环
}
