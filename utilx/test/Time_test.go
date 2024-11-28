package test

import (
	"testing"
	"time"

	"github.com/5-say/go-tool/utilx"
)

func Test_Time(t *testing.T) {
	var (
		Time = utilx.TimeT{
			LocationName: "Asia/Shanghai",
		}
		now = time.Now().Unix()
	)
	t.Log(Time.LocMonthStartDateStr(now))
	t.Log(Time.LocMonthStartTimestamp(now))
	t.Log(Time.LocMonthEndTimestamp(now))
}
