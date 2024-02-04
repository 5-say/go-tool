package rsatool

// 密码学标准
type CryptographyStandards int

const (
	PKCS1 CryptographyStandards = iota + 1
	PKCS8
)
