package rsatool

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 将公钥编码为指定 密码学标准 的 PEM 格式
//
//	publicKey *rsa.PublicKey
//	cs        rsatool.CryptographyStandards 密码学标准
//
// ex:
//
//	rsatool.PEM_EncodePublicKey(publicKey, rsatool.PKCS1)
//	rsatool.PEM_EncodePublicKey(publicKey, rsatool.PKCS8)
func PEM_EncodePublicKey(publicKey *rsa.PublicKey, cs CryptographyStandards) ([]byte, error) {

	if cs == PKCS1 {
		return pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		}), nil
	}

	if cs == PKCS8 {
		block, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return nil, err
		}
		return pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: block,
		}), nil
	}

	return nil, errors.New("unsupported cryptography standards")
}
