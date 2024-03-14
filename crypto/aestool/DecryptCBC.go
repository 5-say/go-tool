package aestool

import (
	"crypto/aes"
	"crypto/cipher"
)

// 以 CBC模式 进行解密
//
//	key        []byte
//	ciphertext []byte
//
// e.g.
//
//	aestool.EncryptCBC(key, ciphertext)
func DecryptCBC(key, ciphertext []byte) (original []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 初始化向量容器
	// 通常为 16 字节（AES算法的块大小）的随机数，和加密时使用的相同。
	// 注意：这个值应该保密，并且在加密和解密过程中必须一致。如果初始向量不正确或被篡改，解密将失败。
	iv := make([]byte, block.BlockSize())

	// 从密文中提取初始向量，将其存储在iv中。
	// 注意：这里使用了 copy 函数来复制初始向量，而不是直接使用 ciphertext 的前 16 个字节，因为我们需要一个独立的副本以便稍后解密。
	// 同时确保 ciphertext 至少包含 16 个字节以避免数组越界错误。
	copy(iv, ciphertext[:block.BlockSize()])

	// 初始化原文容器
	original = make([]byte, len(ciphertext))

	// 创建一个CBC模式下的cipher.Block接口对象，并使用密钥和初始向量进行初始化。注意：这里使用了复制的初始向量而不是直接使用ciphertext的前16个字节，因为我们需要一个独立的副本以便稍后解密。同时确保ciphertext至少包含16个字节以避免数组越界错误。解密过程将从ciphertext中读取每个块并将其解密到decryptedText中。最后返回decryptedText作为明文。注意：由于CBC模式是块加密，解密过程不会返回一个连续的明文字符串。相反，它返回一个包含每个明文块的字节切片。因此，在解密后，您可能需要将每个块重新组合成原始的明文字符串。这可以通过简单地连接decryptedText中的所有块来完成。
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密
	mode.CryptBlocks(original, ciphertext)

	// 块数据解出
	padding := int(original[len(original)-1]) % block.BlockSize()
	original = original[block.BlockSize() : len(original)-padding]

	return
}
