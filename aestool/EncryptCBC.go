package aestool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// 以 CBC模式 进行加密
//
//	key      []byte
//	original []byte
//
// ex:
//
//	aestool.EncryptCBC(key, original)
func EncryptCBC(key, original []byte) (ciphertext []byte, err error) {
	// 创建一个 cipher.Block 接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// 初始化向量
	iv := make([]byte, block.BlockSize())
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}

	// 数据填充
	paddingSize := block.BlockSize() - len(original)%block.BlockSize()
	paddingBytes := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	original = append(iv, original...)
	original = append(original, paddingBytes...)

	// 初始化密文容器
	ciphertext = make([]byte, len(original))

	// 实例化一个块模式，该模式使用给定的块以密码块链接模式进行加密。IV的长度必须与块的块大小相同。
	mode := cipher.NewCBCEncrypter(block, iv)
	// 开始加密
	mode.CryptBlocks(ciphertext, original) // 注意：这里不返回加密后的密文，而是直接将密文存储在 ciphertext 中，因为CBC模式是块加密，不能简单地拼接所有块来得到最终的密文。

	return
}
