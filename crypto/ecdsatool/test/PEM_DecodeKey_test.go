package test

import (
	"testing"

	"github.com/5-say/go-tool/crypto/ecdsatool"
)

var testPrivateKeyPKCS8PEM = []byte(`
-----BEGIN PRIVATE KEY-----
MIHuAgEAMBAGByqGSM49AgEGBSuBBAAjBIHWMIHTAgEBBEIAXymXZxlkGaaR9YF7
DVkSTxbsngiKgBMBPOF2ESpJJPD/Foi33bG0mw2TXqfs88cnaTom2akIn6q7MYCa
t87Jv82hgYkDgYYABAAlPhQU5Ytn9XFILjfk2K9RWtk8hRq0mXEvjNbvde2iElDY
Vpbt4YSsgQBiIO00v26mdwZuv6G7j5jOIHfp3GGr0AGN3gvLiXZEXtqyM+PBimdI
CQLfbsLzAcL51lNy2NXk4tkBDAlyP0QlCEELG3ygF4IBdASACxBujsXIoh/ZjLvT
Cw==
-----END PRIVATE KEY-----`)

var testPublicKeyPKCS8PEM = []byte(`
-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAJT4UFOWLZ/VxSC435NivUVrZPIUa
tJlxL4zW73XtohJQ2FaW7eGErIEAYiDtNL9upncGbr+hu4+YziB36dxhq9ABjd4L
y4l2RF7asjPjwYpnSAkC327C8wHC+dZTctjV5OLZAQwJcj9EJQhBCxt8oBeCAXQE
gAsQbo7FyKIf2Yy70ws=
-----END PUBLIC KEY-----`)

func Test_PEM_DecodeKey(t *testing.T) {

	pk, _ := ecdsatool.PEM_DecodePrivateKey(testPrivateKeyPKCS8PEM)

	pbk, _ := ecdsatool.PEM_DecodePublicKey(testPublicKeyPKCS8PEM)

	if !pk.PublicKey.Equal(pbk) {
		t.Fail()
	}

}
