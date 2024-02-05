package test

import (
	"testing"

	"github.com/5-say/go-tool/gin/jwtx/tool"
	"github.com/golang-jwt/jwt/v5"
)

func Test_GenerateToken(t *testing.T) {
	key, err := tool.GetPrivateKey("jwtx/admin.key")
	// key, err := ecdsatool.GenerateKey(ecdsatool.P_384)
	if err != nil {
		t.Fatal(err)
	}

	token, err := tool.GenerateToken(key, jwt.MapClaims{
		"test": "demo",
	}, tool.P_384)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := tool.ParseToken(token, &key.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims["test"])
}
