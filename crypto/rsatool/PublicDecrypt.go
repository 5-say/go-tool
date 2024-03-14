package rsatool

import (
	"crypto/rsa"

	"github.com/5-say/go-tool/crypto/rsatool/special"
)

// 公钥解密
//
//	publicKey  *rsa.PublicKey
//	ciphertext []byte
//
// e.g.
//
//	rsatool.PublicDecrypt(publicKey, ciphertext)
func PublicDecrypt(publicKey *rsa.PublicKey, ciphertext []byte) (original []byte, err error) {
	return special.PublicDecrypt(publicKey, ciphertext)
}
