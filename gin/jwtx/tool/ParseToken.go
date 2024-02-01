package tool

import (
	"crypto/ecdsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// ParseToken ..
func ParseToken(tokenStr string, publicKey *ecdsa.PublicKey) (claims jwt.MapClaims, err error) {

	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims parse fail")
	}

	return
}
