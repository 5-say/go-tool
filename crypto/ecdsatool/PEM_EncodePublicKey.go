package ecdsatool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

// 将私钥 编码为 PKCS#8 密码学标准 的 PEM 格式
//
//	publicKey *ecdsa.PublicKey
//
// e.g.
//
//	ecdsatool.PEM_EncodePublicKey(publicKey)
func PEM_EncodePublicKey(publicKey *ecdsa.PublicKey) ([]byte, error) {

	block, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: block,
	}), nil

}
