package tool

// 数学曲线标准
type CurveStandard int

const (
	P_256 CurveStandard = iota + 1
	P_384
	P_521
)
