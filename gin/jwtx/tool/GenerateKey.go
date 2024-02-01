package tool

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// generateKey .. 生成椭圆曲线密钥对（转换为 PKCS#8 格式的 PEM 字符串）
func generateKey(c elliptic.Curve) (privateKeyPEMBytes, publicKeyPEMBytes []byte) {
	// 基于椭圆曲线生成私钥
	privateKey, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		fmt.Println("无法生成私钥:", err)
		return
	}

	// 将私钥转换为 PEM 格式的字符串
	PKCS8PrivateKey, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: PKCS8PrivateKey,
	}
	privateKeyPEMBytes = pem.EncodeToMemory(privateKeyPEM)

	// 将公钥转换为 PEM 格式的字符串
	PKIXPublicKey, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	publicKeyPEM := &pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: PKIXPublicKey,
	}
	publicKeyPEMBytes = pem.EncodeToMemory(publicKeyPEM)

	return
}

// GenerateKeyP256 ..
func GenerateKeyP256() (privateKeyPEMBytes, publicKeyPEMBytes []byte) {
	return generateKey(elliptic.P256())
}

// GenerateKeyP384 ..
func GenerateKeyP384() (privateKeyPEMBytes, publicKeyPEMBytes []byte) {
	return generateKey(elliptic.P384())
}

// GenerateKeyP521 ..
func GenerateKeyP521() (privateKeyPEMBytes, publicKeyPEMBytes []byte) {
	return generateKey(elliptic.P521())
}
