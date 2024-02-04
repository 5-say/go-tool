package test

import (
	"fmt"
	"testing"

	"github.com/5-say/go-tool/rsatool"
)

func Test_PEM_EncodeKey(t *testing.T) {
	privateKey, _ := rsatool.GenerateKey(128)

	privateKeyPKCS1PEM, _ := rsatool.PEM_EncodePrivateKey(privateKey, rsatool.PKCS1)
	fmt.Println(string(privateKeyPKCS1PEM))

	privateKeyPKCS8PEM, _ := rsatool.PEM_EncodePrivateKey(privateKey, rsatool.PKCS8)
	fmt.Println(string(privateKeyPKCS8PEM))

	publicKeyPKCS1PEM, _ := rsatool.PEM_EncodePublicKey(&privateKey.PublicKey, rsatool.PKCS1)
	fmt.Println(string(publicKeyPKCS1PEM))

	publicKeyPKCS8PEM, _ := rsatool.PEM_EncodePublicKey(&privateKey.PublicKey, rsatool.PKCS8)
	fmt.Println(string(publicKeyPKCS8PEM))

	// t.Fail()
}
