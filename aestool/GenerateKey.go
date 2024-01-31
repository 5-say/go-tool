package aestool

import (
	"crypto/rand"
)

// GenerateKey .. 生成密钥
//
//	blockSize BlockSize
//
// 16, 24, or 32， AES-128, AES-192, or AES-256.
func GenerateKey(blockSize int) (key []byte, err error) {
	key = make([]byte, blockSize)
	_, err = rand.Read(key)
	return
}
