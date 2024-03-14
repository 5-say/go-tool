package passtool

import (
	"golang.org/x/crypto/bcrypt"
)

// Generate .. 生成
func Generate(password string) (hash string, err error) {
	cost := bcrypt.DefaultCost
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	hash = string(hashByte[:])
	return hash, err
}
