package test

import (
	"fmt"
	"testing"

	"github.com/5-say/go-tool/ecdsatool"
)

func Test_GenerateKey(t *testing.T) {
	key, _ := ecdsatool.GenerateKey(ecdsatool.P_384)
	privateKeyPEMBytes, _ := ecdsatool.PEM_EncodePrivateKey(key)
	fmt.Println()
	fmt.Println("密钥格式: NIST P-384")
	fmt.Printf("\n%s\n", privateKeyPEMBytes)
	fmt.Println("密钥格式: NIST P-384")
	fmt.Println()
}
