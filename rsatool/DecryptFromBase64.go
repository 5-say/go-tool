package rsatool

import "encoding/base64"

// DecryptFromBase64 .. 私钥解密 Base64
func DecryptFromBase64(privateKeyBase64, ciphertextBase64 string) (originalStr string, err error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return
	}

	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return
	}

	originalBytes, err := Decrypt(privateKeyBytes, ciphertextBytes)
	originalStr = string(originalBytes)

	return
}
