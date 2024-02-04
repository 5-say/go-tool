package rsatool

import (
	"crypto/rsa"

	"github.com/5-say/go-tool/rsatool/special"
)

// PrivateEncrypt .. 私钥加密
func PrivateEncrypt(privateKey *rsa.PrivateKey, original []byte) (ciphertext []byte, err error) {
	return special.PrivateEncrypt(privateKey, original)
}
