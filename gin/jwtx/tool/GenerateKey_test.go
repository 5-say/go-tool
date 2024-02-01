package tool

import "testing"

func TestGenerateKey(t *testing.T) {
	privateKeyPEMBytes, publicKeyPEMBytes := GenerateKeyP384()
	t.Log(string(privateKeyPEMBytes), string(publicKeyPEMBytes))
}
