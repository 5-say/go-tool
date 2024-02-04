package rsatool

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func PEM_DecodePrivateKey(pemPrivateKey []byte) (*rsa.PrivateKey, error) {
	// 解析 PEM
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// 解析 私钥
	switch block.Type {
	case "RSA PRIVATE KEY":
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return privateKey.(*rsa.PrivateKey), nil
	}
	return nil, errors.New("unsupported type")
}
