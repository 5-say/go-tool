package rsatool

import (
	"crypto/rsa"

	"github.com/5-say/go-tool/rsatool/special"
)

// 私钥加密
//
//	privateKey *rsa.PrivateKey
//	original   []byte
//
// ex:
//
//	rsatool.PrivateEncrypt(privateKey, original)
func PrivateEncrypt(privateKey *rsa.PrivateKey, original []byte) (ciphertext []byte, err error) {
	return special.PrivateEncrypt(privateKey, original)
}
