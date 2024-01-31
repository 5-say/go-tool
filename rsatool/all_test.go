package rsatool

import (
	"testing"
)

func TestGenerateKeyToBase64(t *testing.T) {
	privateKeyBase64, publicKeyBase64, err := GenerateKeyToBase64(8192)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(privateKeyBase64)
	t.Log(publicKeyBase64)
	// t.Log(EncryptToBase64(publicKeyBase64, "1631498004#9876351"))
}
