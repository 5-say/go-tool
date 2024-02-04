package special

import (
	"crypto/rsa"
)

// PublicDecrypt .. 公钥解密
func PublicDecrypt(publicKey *rsa.PublicKey, ciphertext []byte) (original []byte, err error) {
	var (
		total = len(ciphertext)
		size  = publicKey.Size()
		start = 0
		end   = size
	)

	// 分段解密
	for start < total {
		// 获取指定长度的子串
		if end > total {
			end = total
		}
		sub := ciphertext[start:end]

		// 解密
		subOriginal := RSA_public_decrypt(publicKey, sub)
		// 拼接解密结果
		original = append(original, subOriginal...)

		// 更新起始位置和结束位置
		start += size
		end += size
	}

	return original, nil
}
