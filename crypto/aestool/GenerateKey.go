package aestool

import (
	"crypto/rand"
)

// 生成密钥
//
//	blockSize BlockSize
//
// e.g.
//
//	aestool.GenerateKey(aestool.AES_128)
//	aestool.GenerateKey(aestool.AES_192)
//	aestool.GenerateKey(aestool.AES_256)
func GenerateKey(blockSize BlockSize) (key []byte, err error) {
	key = make([]byte, blockSize)
	_, err = rand.Read(key)
	return
}
