package rsatool

import (
	"crypto/rsa"

	"github.com/5-say/go-tool/rsatool/special"
)

// PublicDecrypt .. 公钥解密
func PublicDecrypt(publicKey *rsa.PublicKey, ciphertext []byte) (original []byte, err error) {
	return special.PublicDecrypt(publicKey, ciphertext)
}
