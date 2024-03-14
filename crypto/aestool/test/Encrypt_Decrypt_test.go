package test

import (
	"encoding/base64"
	"testing"

	"github.com/5-say/go-tool/crypto/aestool"
	"github.com/go-playground/assert/v2"
)

func Test_Encrypt_Decrypt(t *testing.T) {
	var (
		key, _   = aestool.GenerateKey(aestool.AES_256)
		original = []byte(`145678987654456`)
	)

	ciphertext, err := aestool.EncryptCBC(key, original)
	t.Log(base64.StdEncoding.EncodeToString(key))
	t.Log(base64.StdEncoding.EncodeToString(ciphertext))
	assert.Equal(t, err, nil)

	o, err := aestool.DecryptCBC(key, ciphertext)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(o), string(original))

	// t.Fail()
}
