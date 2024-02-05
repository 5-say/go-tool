package test

import (
	"encoding/base64"
	"testing"

	"github.com/5-say/go-tool/aestool"
)

func Test_Encrypt_Decrypt(t *testing.T) {
	var (
		key, _   = aestool.GenerateKey(aestool.AES_128)
		original = []byte(`demo`)
	)
	ciphertext, err := aestool.EncryptCBC(key, original)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(ciphertext))

	o, err := aestool.DecryptCBC(key, ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(o))

	if string(o) != string(original) {
		t.Fail()
	}
}
