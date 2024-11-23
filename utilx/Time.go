package utilx

import "time"

type TimeT struct {
	LocationName string // "Asia/Shanghai"
}

func (s TimeT) Loc() *time.Location {
	var loc, _ = time.LoadLocation(s.LocationName)
	return loc
}

// 获取目标时间戳所在日，起始日期字符串（自动修正时区误差）
//
//	unixTimestamp int64 // UTC时间戳
func (s TimeT) LocDayStartDateStr(unixTimestamp int64) string {
	todayStartTime, _ := time.ParseInLocation("2006-01-02", time.Unix(unixTimestamp, 0).In(s.Loc()).Format("2006-01-02"), s.Loc())
	return todayStartTime.Format("2006-01-02")
}

// 获取目标时间戳所在日，起始时间戳（自动修正时区误差）
//
//	unixTimestamp int64 // UTC时间戳
func (s TimeT) LocDayStartTimestamp(unixTimestamp int64) int64 {
	todayStartTime, _ := time.ParseInLocation("2006-01-02", time.Unix(unixTimestamp, 0).In(s.Loc()).Format("2006-01-02"), s.Loc())
	return todayStartTime.Unix()
}

// 获取目标时间戳所在日，结束时间戳（自动修正时区误差）
//
//	unixTimestamp int64 // UTC时间戳
func (s TimeT) LocDayEndTimestamp(unixTimestamp int64) int64 {
	todayStartTime, _ := time.ParseInLocation("2006-01-02", time.Unix(unixTimestamp, 0).In(s.Loc()).Format("2006-01-02"), s.Loc())
	return todayStartTime.Add(24*time.Hour).Unix() - 1
}
