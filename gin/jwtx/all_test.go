package jwtx

import (
	"fmt"
	"testing"

	"github.com/5-say/go-tool/gin/jwtx/tool"
)

func Test_GenerateKeyP384(t *testing.T) {
	privateKeyPEMBytes, _ := tool.GenerateKeyP384()
	fmt.Println()
	fmt.Println("密钥格式: NIST P-384")
	fmt.Printf("\n%s\n", privateKeyPEMBytes)
	fmt.Println("密钥格式: NIST P-384")
	fmt.Println()
}
