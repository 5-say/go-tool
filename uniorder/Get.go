package uniorder

// 取得分组唯一订单号，并发安全
//
// e.g.
//
//	cycleNum, inCycleIndex := uniorder.Get("demo")
func Get(group string) (cycleNum, inCycleIndex uint64) {
	var g = uniBox[group]
	var cache = <-g.UniCache
	return cache.CycleNum, cache.InCycleIndex
}
