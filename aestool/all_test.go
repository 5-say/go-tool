package aestool

import "testing"

func Test_GenerateKeyToBase64(t *testing.T) {
	keyBase64, err := GenerateKeyToBase64(32)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(keyBase64)
	// t.Fail()
}

func TestXxx(t *testing.T) {
	keyBase64 := "U+ydWocmE7aIonFxF8aY5a+bt/7m3otTo7N/U6q/Z64="
	ciphertextBase64, err := EncryptCBCToBase64(keyBase64, "demo567823456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ciphertextBase64)
	originalStr, err := DecryptCBCFromBase64(keyBase64, ciphertextBase64)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(originalStr)
	// t.Fail()
}
