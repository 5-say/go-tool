package test

import (
	"encoding/base64"
	"testing"

	"github.com/5-say/go-tool/crypto/rsatool"
)

func Test_Encrypt_Decrypt(t *testing.T) {
	var (
		privateKey, _ = rsatool.GenerateKey(128)
		publicKey     = &privateKey.PublicKey
		original      = []byte(`g943hg943hg9344g943hg943hg9344g943hg943hg9344g943hg943hg9344g943hg943hg9344g943hg943hg9344g943hg943hg9344g943hg943hg934`)
	)
	ciphertext, err := rsatool.Encrypt(publicKey, original)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(ciphertext))

	o, err := rsatool.Decrypt(privateKey, ciphertext)
	if err != nil {
		t.Fatal(err)
	}

	if string(o) != string(original) {
		t.Log(string(o))
		t.Log(string(original))
		t.Fail()
	}
}
