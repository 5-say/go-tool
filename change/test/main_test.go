package test

import (
	"testing"

	"github.com/5-say/go-tool/change"
	"github.com/go-playground/assert/v2"
)

func Test_Main(t *testing.T) {
	var x float32 = 67.8
	var a *float32
	_, err := change.To[uint8](a)
	assert.Equal(t, err, change.ErrNilPointer)
	aaa := change.ForceTo[uint8](a, 89)
	assert.Equal(t, aaa, uint8(89))
	a = &x
	aaaa := change.ForceTo[uint8](a, 89)
	assert.Equal(t, aaaa, uint8(67))
}
