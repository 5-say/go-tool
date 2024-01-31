package rsatool

import "encoding/base64"

// GenerateKeyToBase64 .. 生成公钥私钥 Base64
// bits 加密强度：1024、2048、3072、4096、8192（推荐使用 2048）
func GenerateKeyToBase64(bits int) (privateKeyBase64, publicKeyBase64 string, err error) {
	privateKeyBytes, publicKeyBytes, err := GenerateKey(bits)
	if err != nil {
		return
	}

	privateKeyBase64 = base64.StdEncoding.EncodeToString(privateKeyBytes)
	publicKeyBase64 = base64.StdEncoding.EncodeToString(publicKeyBytes)

	return
}
