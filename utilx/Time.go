package utilx

import (
	"time"
)

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

// 获取目标时间戳所在月，起始日期字符串（自动修正时区误差）
//
//	unixTimestamp int64 // UTC时间戳
func (s TimeT) LocMonthStartDateStr(unixTimestamp int64) string {
	thisMonthStartTime, _ := time.ParseInLocation("2006-01", time.Unix(unixTimestamp, 0).In(s.Loc()).Format("2006-01"), s.Loc())
	return thisMonthStartTime.Format("2006-01")
}

// 获取目标时间戳所在日，起始时间戳（自动修正时区误差）
//
//	unixTimestamp int64 // UTC时间戳
func (s TimeT) LocMonthStartTimestamp(unixTimestamp int64) int64 {
	thisMonthStartTime, _ := time.ParseInLocation("2006-01", time.Unix(unixTimestamp, 0).In(s.Loc()).Format("2006-01"), s.Loc())
	return thisMonthStartTime.Unix()
}

// 获取目标时间戳所在日，结束时间戳（自动修正时区误差）
//
//	unixTimestamp int64 // UTC时间戳
func (s TimeT) LocMonthEndTimestamp(unixTimestamp int64) int64 {
	var (
		locTime = time.Unix(unixTimestamp, 0).In(s.Loc())
		year    = locTime.Year()
		month   = locTime.Month() + 1
	)
	if month > 12 {
		year += 1
		month = 1
	}
	return time.Date(year, month, 1, 0, 0, 0, 0, s.Loc()).Unix() - 1
}
