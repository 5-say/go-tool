package aestool

import "encoding/base64"

// EncryptCBCToBase64 .. 加密 Base64
func EncryptCBCToBase64(keyBase64, originalStr string) (ciphertextBase64 string, err error) {
	keyBytes, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return
	}

	ciphertext, err := EncryptCBC(keyBytes, []byte(originalStr))
	ciphertextBase64 = base64.StdEncoding.EncodeToString(ciphertext)

	return
}
