package test

import (
	"crypto"
	"encoding/base64"
	"testing"

	"github.com/5-say/go-tool/rsatool"
)

func Test_Sign_Verify(t *testing.T) {
	var (
		privateKey, _ = rsatool.GenerateKey(128)
		original      = []byte(`demo`)
	)
	signature, err := rsatool.Sign(privateKey, original, crypto.SHA512)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(signature))

	err = rsatool.Verify(&privateKey.PublicKey, original, signature, crypto.SHA512)
	if err != nil {
		t.Fatal(err)
	}
}
