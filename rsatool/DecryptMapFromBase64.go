package rsatool

import "encoding/json"

// DecryptMapFromBase64 .. 私钥解密 Base64
func DecryptMapFromBase64(privateKeyBase64, ciphertextBase64 string) (originalMap map[string]interface{}, err error) {
	originalStr, err := DecryptFromBase64(privateKeyBase64, ciphertextBase64)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(originalStr), &originalMap)
	return
}
