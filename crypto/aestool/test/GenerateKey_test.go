package test

import (
	"testing"

	"github.com/5-say/go-tool/crypto/aestool"
	"github.com/go-playground/assert/v2"
)

func Test_Base64_GenerateKey(t *testing.T) {
	key, err := aestool.GenerateKey(aestool.AES_256)
	assert.Equal(t, err, nil)
	t.Log(key)
	t.Fail()
}
