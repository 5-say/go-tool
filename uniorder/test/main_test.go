package test

import (
	"testing"
	"time"

	"github.com/5-say/go-tool/uniorder"
)

func TestXxx(t *testing.T) {
	uniorder.OriginTime = time.Date(2013, 4, 5, 0, 0, 0, 0, time.UTC)
	uniorder.InitGroup("demo", 18, 0, 0)
	go func() {
		for {
			time.Sleep(time.Millisecond * 9)
			t.Log(uniorder.Get("demo"))
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 13)
			t.Log(uniorder.Get("demo"))
		}
	}()
	time.Sleep(time.Second * 1)
	// t.Fail()
}
