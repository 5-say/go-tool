package tool

import (
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken ..
func GenerateToken(key *ecdsa.PrivateKey, claims jwt.MapClaims, cs CurveStandard) (token string, err error) {
	switch cs {
	case P_256:
		return jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(key)
	case P_384:
		return jwt.NewWithClaims(jwt.SigningMethodES384, claims).SignedString(key)
	case P_521:
		return jwt.NewWithClaims(jwt.SigningMethodES512, claims).SignedString(key)
	default:
		return "", nil
	}
}
