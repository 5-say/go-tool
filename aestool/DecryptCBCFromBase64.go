package aestool

import "encoding/base64"

// DecryptCBCFromBase64 .. 解密 Base64
func DecryptCBCFromBase64(keyBase64, ciphertextBase64 string) (originalStr string, err error) {
	keyBytes, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return
	}

	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return
	}

	originalBytes, err := DecryptCBC(keyBytes, ciphertextBytes)
	originalStr = string(originalBytes)

	return
}
