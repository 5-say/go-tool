package rsatool

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"

	"golang.org/x/crypto/sha3"
)

// 私钥签名
//
//	privateKey *rsa.PrivateKey
//	original   []byte
//	hashFunc   crypto.Hash
//
// ex:
//
//	rsatool.Sign(privateKey, original, crypto.SHA256)
func Sign(privateKey *rsa.PrivateKey, original []byte, hashFunc crypto.Hash) (signature []byte, err error) {
	var h hash.Hash
	switch hashFunc {
	case crypto.MD5:
		h = md5.New()
	case crypto.SHA1:
		h = sha1.New()
	case crypto.SHA224:
		h = sha256.New224()
	case crypto.SHA256:
		h = sha256.New()
	case crypto.SHA384:
		h = sha512.New384()
	case crypto.SHA512:
		h = sha512.New()
	case crypto.SHA3_224:
		h = sha3.New224()
	case crypto.SHA3_256:
		h = sha3.New256()
	case crypto.SHA3_384:
		h = sha3.New384()
	case crypto.SHA3_512:
		h = sha3.New512()
	case crypto.SHA512_224:
		h = sha512.New512_224()
	case crypto.SHA512_256:
		h = sha512.New512_256()
	default:
		return nil, errors.New("unsupported hash function")
	}
	h.Write(original)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, hashFunc, hashed)
}
