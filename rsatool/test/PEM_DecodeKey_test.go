package test

import (
	"testing"

	"github.com/5-say/go-tool/rsatool"
)

var testPrivateKeyPKCS1PEM = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDqQ0vXtFVP1276HLLg8VJWg056kKA8PdWgNhLsPXJrSM3GKJ3p
ZwxGA7EDvhsWMHbwI4qYgHFURGv+dNx4sgq1aFdlJlQKgZ5i61/01G2mEMWzXbyM
QT0RvQLMdPubcZ9zotY53K3A/W4G1hMY/Zq02jz/2/Io1pCKbJkceRXnIwIDAQAB
AoGAPPpFbKFpL34xSfNccIcAWrqVG/w3aVbjG2/X3xxjgx+RSIpcCFwlargdRI8g
d9cnrnxh4C5W0Yt53ONIXDZj6LMBlfRBDVWw0YGn44waI8SyKBNXSyb4EGLDlkhy
KYRWPoSIrxhDzbFE5QZjCP4TdAMQb5HsZKNNneQi8hf0H2kCQQD18YPmvNaaSgkB
ic5VRwuGymGHOHD7K1ZIR/cW4H9SB+K165Rc2WsIRQIkQl/QaLd5UN04JDiL4FEl
GZdnYqL1AkEA89eC8s0Rb9+gNoACNeSAPBFeNgdF52dLnQIhOecoNhsShumpXzD5
GEg+Mjoj18FVNI5GDOWvllVikQiaWaeCtwJBAKO+TXQky7hjjEVy62XJbDosqewu
PsJq4wIyr6aShxIlX+KA2wFKW8FguIFNk58c6PKjTPWg3+j4yu2pcS0K9bUCQQDX
qzyx3bjXdzYwHgZWflUBS+dE6Cfm5NDW/ORAX3iMamN3mkFo5VrcJEGGkS/Ui7X9
9vx0Oc6Uq86zMqW6OFN/AkAY2oN9EGJvIJNBB9Kj2yw+yhdpSdiXbT/PpZfnRosK
E0clmrL7KcWiThhTYZ4fPTdpax/joJ4yjl8ME4fpb/je
-----END RSA PRIVATE KEY-----`)

var testPrivateKeyPKCS8PEM = []byte(`
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAOpDS9e0VU/Xbvoc
suDxUlaDTnqQoDw91aA2Euw9cmtIzcYonelnDEYDsQO+GxYwdvAjipiAcVREa/50
3HiyCrVoV2UmVAqBnmLrX/TUbaYQxbNdvIxBPRG9Asx0+5txn3Oi1jncrcD9bgbW
Exj9mrTaPP/b8ijWkIpsmRx5FecjAgMBAAECgYA8+kVsoWkvfjFJ81xwhwBaupUb
/DdpVuMbb9ffHGODH5FIilwIXCVquB1EjyB31yeufGHgLlbRi3nc40hcNmPoswGV
9EENVbDRgafjjBojxLIoE1dLJvgQYsOWSHIphFY+hIivGEPNsUTlBmMI/hN0AxBv
kexko02d5CLyF/QfaQJBAPXxg+a81ppKCQGJzlVHC4bKYYc4cPsrVkhH9xbgf1IH
4rXrlFzZawhFAiRCX9Bot3lQ3TgkOIvgUSUZl2diovUCQQDz14LyzRFv36A2gAI1
5IA8EV42B0XnZ0udAiE55yg2GxKG6alfMPkYSD4yOiPXwVU0jkYM5a+WVWKRCJpZ
p4K3AkEAo75NdCTLuGOMRXLrZclsOiyp7C4+wmrjAjKvppKHEiVf4oDbAUpbwWC4
gU2Tnxzo8qNM9aDf6PjK7alxLQr1tQJBANerPLHduNd3NjAeBlZ+VQFL50ToJ+bk
0Nb85EBfeIxqY3eaQWjlWtwkQYaRL9SLtf32/HQ5zpSrzrMypbo4U38CQBjag30Q
Ym8gk0EH0qPbLD7KF2lJ2JdtP8+ll+dGiwoTRyWasvspxaJOGFNhnh89N2lrH+Og
njKOXwwTh+lv+N4=
-----END PRIVATE KEY-----`)

var testPublicKeyPKCS1PEM = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAOpDS9e0VU/XbvocsuDxUlaDTnqQoDw91aA2Euw9cmtIzcYonelnDEYD
sQO+GxYwdvAjipiAcVREa/503HiyCrVoV2UmVAqBnmLrX/TUbaYQxbNdvIxBPRG9
Asx0+5txn3Oi1jncrcD9bgbWExj9mrTaPP/b8ijWkIpsmRx5FecjAgMBAAE=
-----END RSA PUBLIC KEY-----
`)

var testPublicKeyPKCS8PEM = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDqQ0vXtFVP1276HLLg8VJWg056
kKA8PdWgNhLsPXJrSM3GKJ3pZwxGA7EDvhsWMHbwI4qYgHFURGv+dNx4sgq1aFdl
JlQKgZ5i61/01G2mEMWzXbyMQT0RvQLMdPubcZ9zotY53K3A/W4G1hMY/Zq02jz/
2/Io1pCKbJkceRXnIwIDAQAB
-----END PUBLIC KEY-----
`)

func Test_PEM_DecodeKey(t *testing.T) {

	pk1, _ := rsatool.PEM_DecodePrivateKey(testPrivateKeyPKCS1PEM)
	pk2, _ := rsatool.PEM_DecodePrivateKey(testPrivateKeyPKCS8PEM)

	if !pk1.Equal(pk2) {
		t.Fail()
	}

	pbk1, _ := rsatool.PEM_DecodePublicKey(testPublicKeyPKCS1PEM)
	pbk2, _ := rsatool.PEM_DecodePublicKey(testPublicKeyPKCS8PEM)

	if !pbk1.Equal(pbk2) {
		t.Fail()
	}

}
