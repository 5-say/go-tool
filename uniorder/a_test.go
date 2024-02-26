package uniorder

import (
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	OriginTime = time.Date(2013, 4, 5, 0, 0, 0, 0, time.UTC)
	InitGroup("demo", 18, 0, 0)
	go func() {
		for {
			time.Sleep(time.Millisecond * 9)
			t.Log(Get("demo"))
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 13)
			t.Log(Get("demo"))
		}
	}()
	time.Sleep(time.Second * 1)
	t.Fail()
}
