package aestool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// EncryptCBC .. 加密
func EncryptCBC(keyBytes, originalBytes []byte) (ciphertextBytes []byte, err error) {
	// 创建一个 cipher.Block 接口对象
	block, err := aes.NewCipher(keyBytes)
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
	paddingSize := block.BlockSize() - len(originalBytes)%block.BlockSize()
	paddingBytes := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	originalBytes = append(iv, originalBytes...)
	originalBytes = append(originalBytes, paddingBytes...)

	// 初始化密文容器
	ciphertextBytes = make([]byte, len(originalBytes))

	// 实例化一个块模式，该模式使用给定的块以密码块链接模式进行加密。IV的长度必须与块的块大小相同。
	mode := cipher.NewCBCEncrypter(block, iv)
	// 开始加密
	mode.CryptBlocks(ciphertextBytes, originalBytes) // 注意：这里不返回加密后的密文，而是直接将密文存储在 ciphertextBytes 中，因为CBC模式是块加密，不能简单地拼接所有块来得到最终的密文。

	return
}
