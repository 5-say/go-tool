package rsatool

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解码 PEM 格式 公钥
//
//	pemPublicKey []byte
//
// ex:
//
//	rsatool.PEM_DecodePublicKey(pemPublicKey)
func PEM_DecodePublicKey(pemPublicKey []byte) (*rsa.PublicKey, error) {
	// 解析 PEM
	block, _ := pem.Decode(pemPublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析 私钥
	switch block.Type {
	case "RSA PUBLIC KEY":
		return x509.ParsePKCS1PublicKey(block.Bytes)
	case "PUBLIC KEY":
		publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return publicKey.(*rsa.PublicKey), nil
	}
	return nil, errors.New("unsupported type")
}
