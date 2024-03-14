package ecdsatool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解码 PEM 格式 公钥
//
//	pemPublicKey []byte
//
// e.g.
//
//	ecdsatool.PEM_DecodePublicKey(pemPublicKey)
func PEM_DecodePublicKey(pemPublicKey []byte) (*ecdsa.PublicKey, error) {
	// 解析 PEM
	block, _ := pem.Decode(pemPublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析 公钥
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return publicKey.(*ecdsa.PublicKey), nil
}
