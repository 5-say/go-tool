package passtool

import (
	"golang.org/x/crypto/bcrypt"
)

// Compare .. 比较
func Compare(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
