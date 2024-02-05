package tool

import (
	"crypto/ecdsa"
	"os"

	"github.com/5-say/go-tool/ecdsatool"
)

// GetPrivateKey .. 从 PEM 格式文件 获取 私钥对象
func GetPrivateKey(filePath string) (privateKey *ecdsa.PrivateKey, err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	return ecdsatool.PEM_DecodePrivateKey(data)
}
