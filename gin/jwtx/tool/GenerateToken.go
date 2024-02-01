package tool

import (
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTokenES256 ..
func GenerateTokenES256(key *ecdsa.PrivateKey, claims jwt.MapClaims) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(key)
}

// GenerateTokenES384 ..
func GenerateTokenES384(key *ecdsa.PrivateKey, claims jwt.MapClaims) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodES384, claims).SignedString(key)
}

// GenerateTokenES512 ..
func GenerateTokenES512(key *ecdsa.PrivateKey, claims jwt.MapClaims) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodES512, claims).SignedString(key)
}
