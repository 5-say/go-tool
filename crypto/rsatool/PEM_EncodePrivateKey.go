package rsatool

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 将私钥编码为指定 密码学标准 的 PEM 格式
//
//	privateKey *rsa.PrivateKey
//	cs         rsatool.CryptographyStandards 密码学标准
//
// e.g.
//
//	rsatool.PEM_EncodePrivateKey(privateKey, rsatool.PKCS1)
//	rsatool.PEM_EncodePrivateKey(privateKey, rsatool.PKCS8)
func PEM_EncodePrivateKey(privateKey *rsa.PrivateKey, cs CryptographyStandards) ([]byte, error) {

	if cs == PKCS1 {
		return pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}), nil
	}

	if cs == PKCS8 {
		block, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, err
		}
		return pem.EncodeToMemory(&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: block,
		}), nil
	}

	return nil, errors.New("unsupported cryptography standards")
}
