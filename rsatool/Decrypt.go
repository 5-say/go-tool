package rsatool

import (
	"crypto/rand"
	"crypto/rsa"
)

// 私钥解密
//
//	privateKey *rsa.PrivateKey
//	ciphertext []byte
//
// ex:
//
//	rsatool.Decrypt(privateKey, ciphertext)
func Decrypt(privateKey *rsa.PrivateKey, ciphertext []byte) (original []byte, err error) {
	var (
		total = len(ciphertext)
		size  = privateKey.Size()
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
		subOriginal, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, sub)
		if err != nil {
			return nil, err
		}
		// 拼接解密结果
		original = append(original, subOriginal...)

		// 更新起始位置和结束位置
		start += size
		end += size
	}

	return original, nil
}
