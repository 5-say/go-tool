package special

import (
	"crypto"
	"crypto/rsa"
)

// PrivateEncrypt .. 私钥加密
func PrivateEncrypt(privateKey *rsa.PrivateKey, original []byte) (ciphertext []byte, err error) {
	var (
		total = len(original)
		size  = privateKey.Size() - 11
		start = 0
		end   = size
	)

	// 分段加密
	for start < total {
		// 获取指定长度的子串
		if end > total {
			end = total
		}
		sub := original[start:end]

		// 加密
		subCiphertext, err := rsa.SignPKCS1v15(nil, privateKey, crypto.Hash(0), sub)
		if err != nil {
			return nil, err
		}
		// 拼接加密结果
		ciphertext = append(ciphertext, subCiphertext...)

		// 更新起始位置和结束位置
		start += size
		end += size
	}

	return ciphertext, nil
}
