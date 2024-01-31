package rsatool

import "encoding/base64"

// EncryptToBase64 .. 公钥加密 Base64
func EncryptToBase64(publicKeyBase64, originalStr string) (ciphertextBase64 string, err error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return
	}

	ciphertext, err := Encrypt(publicKeyBytes, []byte(originalStr))
	ciphertextBase64 = base64.StdEncoding.EncodeToString(ciphertext)

	return
}
