package rsatool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Decrypt .. 私钥解密
func Decrypt(privateKeyBytes, ciphertextBytes []byte) (originalBytes []byte, err error) {
	// 解码私钥
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		err = errors.New("private key error")
		return
	}

	// 解析 PKCS1 格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertextBytes)
}
