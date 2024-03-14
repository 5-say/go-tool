package ecdsatool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解码 PEM 格式 私钥
//
//	pemPrivateKey []byte
//
// e.g.
//
//	ecdsatool.PEM_DecodePrivateKey(pemPrivateKey)
func PEM_DecodePrivateKey(pemPrivateKey []byte) (*ecdsa.PrivateKey, error) {
	// 解析 PEM
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	// 解析 私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey.(*ecdsa.PrivateKey), nil
}
