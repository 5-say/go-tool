package rsatool

import (
	"crypto/rand"
	"crypto/rsa"
)

// GenerateKey .. 生成公钥私钥
//
//	size int modulus size in bytes（128、256、384、512）
func GenerateKey(size int) (privateKey *rsa.PrivateKey, err error) {
	return rsa.GenerateKey(rand.Reader, size*8)
}
