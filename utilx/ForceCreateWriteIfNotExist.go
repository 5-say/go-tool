package utilx

import (
	"io/fs"
	"os"
	"regexp"
)

// 如果不存在则强制创建写入
func ForceCreateWriteIfNotExist(filePath string, perm fs.FileMode, content string) {
	// 解析目录
	dirPath := regexp.MustCompile(`[^\/]+$`).ReplaceAllString(filePath, "")

	// 尝试创建多级目录
	err := os.MkdirAll(dirPath, perm) // 第二个参数是目录的权限
	if err != nil {
		panic(err)
	}

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，则创建文件并写入字符串
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// 写入字符串
		_, err = file.WriteString(content)
		if err != nil {
			panic(err)
		}
	}
}
