package rsatool

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"

	"golang.org/x/crypto/sha3"
)

// hashed[:] 会创建一个新的切片。
// 这个新切片是原切片的浅拷贝，它会复制原切片的元素，但不会复制底层数组。
// 对其中一个切片进行修改不会影响另一个切片。

// 公钥验签
//
//	publicKey *rsa.PublicKey
//	original  []byte
//	signature []byte
//	hashFunc  crypto.Hash
//
// e.g.
//
//	rsatool.Verify(publicKey, original, signature, crypto.SHA256)
func Verify(publicKey *rsa.PublicKey, original, signature []byte, hashFunc crypto.Hash) error {
	switch hashFunc {
	case crypto.MD5:
		hashed := md5.Sum(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA1:
		hashed := sha1.Sum(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA224:
		hashed := sha256.Sum224(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA256:
		hashed := sha256.Sum256(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA384:
		hashed := sha512.Sum384(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA512:
		hashed := sha512.Sum512(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA3_224:
		hashed := sha3.Sum224(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA3_256:
		hashed := sha3.Sum256(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA3_384:
		hashed := sha3.Sum384(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA3_512:
		hashed := sha3.Sum512(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA512_224:
		hashed := sha512.Sum512_224(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	case crypto.SHA512_256:
		hashed := sha512.Sum512_256(original)
		return rsa.VerifyPKCS1v15(publicKey, hashFunc, hashed[:], signature)
	default:
		return errors.New("unsupported hash function")
	}
}
