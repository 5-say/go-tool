package uniorder

import "time"

// 时间原点
var OriginTime = time.Date(2024, 1, 8, 8, 0, 0, 0, time.UTC)

// uniBox ..
var uniBox = map[string]uniGroup{}

// uniGroup ..
type uniGroup = struct {
	LastCycleNum     uint64 // 周期数
	LastInCycleIndex uint64 // 周期内索引

	UniCache chan uniCache
}

// uniCache ..
type uniCache = struct {
	CycleNum     uint64 // 周期数
	InCycleIndex uint64 // 周期内索引
}
