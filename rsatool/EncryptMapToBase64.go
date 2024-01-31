package rsatool

import "encoding/json"

// EncryptMapToBase64 .. 公钥加密 Base64
func EncryptMapToBase64(publicKeyBase64 string, originalMap map[string]interface{}) (ciphertextBase64 string, err error) {
	originalByte, err := json.Marshal(originalMap)
	if err != nil {
		return
	}
	return EncryptToBase64(publicKeyBase64, string(originalByte))
}
