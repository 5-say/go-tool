package rsatool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// GenerateKey .. 生成公钥私钥
// bits 加密强度：1024、2048、3072、4096、8192（推荐使用 2048）
func GenerateKey(bits int) (privateKeyBytes, publicKeyBytes []byte, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateKeyBytes = pem.EncodeToMemory(block)

	// 根据私钥生成公钥
	publicKey := &privateKey.PublicKey
	derBytes := x509.MarshalPKCS1PublicKey(publicKey)
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derBytes,
	}
	publicKeyBytes = pem.EncodeToMemory(block)

	return
}
