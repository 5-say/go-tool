package utilx

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

var Md5 Md5T

type Md5T struct{}

// 字符串 md5
func (Md5T) String(str string) (md5String string) {
	// 创建一个新的hash.Hash对象，用于MD5加密
	hasher := md5.New()

	// 将字符串写入hasher对象
	io.WriteString(hasher, str)

	// 计算MD5值，返回一个字节切片
	hash := hasher.Sum(nil)

	// 将字节切片转换为十六进制字符串
	md5String = hex.EncodeToString(hash)

	return
}
