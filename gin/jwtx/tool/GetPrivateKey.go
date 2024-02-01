package tool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"errors"
)

// GetPrivateKey .. 从 PEM 格式二进制编码 获取 私钥对象
func GetPrivateKey(filePath string) (privateKey *ecdsa.PrivateKey, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	// 解码PEM数据
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	// 解析椭圆曲线密钥
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	// 类型转换
	privateKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("the privateKey is not ecdsa.PrivateKey")
	}

	return
}
