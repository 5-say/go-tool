package ecdsatool

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// 生成椭圆曲线私钥
//
//	curveStandard CurveStandard
//
// e.g.
//
//	ecdsatool.GenerateKey(ecdsatool.P_384)
func GenerateKey(curveStandard CurveStandard) (privateKey *ecdsa.PrivateKey, err error) {
	var curve elliptic.Curve

	switch curveStandard {
	case P_224:
		curve = elliptic.P224()
	case P_256:
		curve = elliptic.P256()
	case P_384:
		curve = elliptic.P384()
	case P_521:
		curve = elliptic.P521()
	}

	// 基于椭圆曲线生成私钥
	return ecdsa.GenerateKey(curve, rand.Reader)
}
