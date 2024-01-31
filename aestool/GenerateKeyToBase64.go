package aestool

import "encoding/base64"

// GenerateKeyToBase64 .. 生成密钥 Base64
//
//	blockSize BlockSize
//
// 16, 24, or 32， AES-128, AES-192, or AES-256.
func GenerateKeyToBase64(blockSize int) (keyBase64 string, err error) {
	keyBytes, err := GenerateKey(blockSize)
	if err != nil {
		return
	}

	keyBase64 = base64.StdEncoding.EncodeToString(keyBytes)

	return
}
