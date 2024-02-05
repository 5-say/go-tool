package test

import (
	"fmt"
	"testing"

	"github.com/5-say/go-tool/ecdsatool"
)

func Test_PEM_EncodeKey(t *testing.T) {
	privateKey, _ := ecdsatool.GenerateKey(ecdsatool.P_521)

	privateKeyPKCS8PEM, _ := ecdsatool.PEM_EncodePrivateKey(privateKey)
	fmt.Println(string(privateKeyPKCS8PEM))

	publicKeyPKCS8PEM, _ := ecdsatool.PEM_EncodePublicKey(&privateKey.PublicKey)
	fmt.Println(string(publicKeyPKCS8PEM))

	// t.Fail()
}
