package ecdsatool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

// 将私钥 编码为 PKCS#8 密码学标准 的 PEM 格式
//
//	privateKey *ecdsa.PrivateKey
//
// e.g.
//
//	ecdsatool.PEM_EncodePrivateKey(privateKey)
func PEM_EncodePrivateKey(privateKey *ecdsa.PrivateKey) ([]byte, error) {

	block, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: block,
	}), nil

}
