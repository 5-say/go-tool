package rsatool

import (
	"crypto/rsa"

	"github.com/5-say/go-tool/rsatool/special"
)

// 公钥解密
//
//	publicKey  *rsa.PublicKey
//	ciphertext []byte
//
// ex:
//
//	rsatool.PublicDecrypt(publicKey, ciphertext)
func PublicDecrypt(publicKey *rsa.PublicKey, ciphertext []byte) (original []byte, err error) {
	return special.PublicDecrypt(publicKey, ciphertext)
}
