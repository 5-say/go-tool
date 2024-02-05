package ecdsatool

// 数学曲线标准
type CurveStandard int

const (
	P_224 CurveStandard = iota + 1
	P_256
	P_384
	P_521
)
