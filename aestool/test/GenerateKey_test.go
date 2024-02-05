package test

import (
	"encoding/base64"
	"testing"

	"github.com/5-say/go-tool/aestool"
)

func Test_Base64_GenerateKey(t *testing.T) {
	key, err := aestool.GenerateKey(aestool.AES_128)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(key))
}
