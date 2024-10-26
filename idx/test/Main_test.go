package test

import (
	"testing"

	"github.com/5-say/go-tool/idx"
)

func TestMain(t *testing.T) {
	var OrderNo = idx.TootT{
		MinLenth:  2,
		Base:      18,
		Offset:    888,
		FillChars: []rune{'x', 'y', 'z'},
		MixMap: map[rune]rune{
			'4': '8',
			'8': 'd',
			'd': '4',
		},
	}
	str := OrderNo.Encode(1287382)
	t.Log(str)
	t.Log(OrderNo.Decode(str))
}
