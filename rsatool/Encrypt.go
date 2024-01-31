package rsatool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Encrypt .. 公钥加密
func Encrypt(publicKeyBytes, originalBytes []byte) (ciphertextBytes []byte, err error) {
	// 解码公钥
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		err = errors.New("public key error")
		return
	}

	// 解析公钥
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return
	}

	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originalBytes)
}
