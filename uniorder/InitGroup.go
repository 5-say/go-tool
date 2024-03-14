package uniorder

import (
	"math/rand"
	"time"
)

// 初始化分组，开启异步生成队列
//
//	group            string // 分组名称
//	rollingHours     uint64 // 滚动周期（小时）
//	lastCycleNum     uint64 // 已使用的最后一个订单号的 周期数
//	lastInCycleIndex uint64 // 已使用的最后一个订单号的 周期内索引
//
// e.g.
//
//	uniorder.InitGroup("demo", 18, 0, 0)
func InitGroup(group string, rollingHours, lastCycleNum, lastInCycleIndex uint64) {
	// 初始化
	uniBox[group] = uniGroup{
		LastCycleNum:     lastCycleNum,
		LastInCycleIndex: lastInCycleIndex,

		// 无缓冲通道
		UniCache: make(chan uniCache),
	}
	var g = uniBox[group]

	// 开启协程
	go func() {
		// 为每个 goroutine 创建一个新的 随机数源 和 生成器
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		for {
			// 计算滚动周期数
			// （该周期小于程序中断后重启的时间，即可确保滚动至新周期）
			// 秒为周期
			cycleNum := uint64(time.Since(OriginTime).Hours()) / rollingHours
			// 同步最新周期，重置索引
			if cycleNum > g.LastCycleNum {
				g.LastCycleNum = cycleNum
				g.LastInCycleIndex = 0
			}
			// 同步最新索引
			g.LastInCycleIndex += uint64(random.Intn(8)) + 1
			// 缓存唯一订单信息
			g.UniCache <- uniCache{
				CycleNum:     g.LastCycleNum,
				InCycleIndex: g.LastInCycleIndex,
			}
		}
	}()
}
