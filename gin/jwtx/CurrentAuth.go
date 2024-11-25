package jwtx

import "github.com/gin-gonic/gin"

type Current struct {
	AccountID     uint64 // 账户 ID
	LoginGroup    string // 登录的分组
	LoginTerminal string // 登录的终端
	MakeTokenIP   string // 首次请求生成 token 的 IP 地址
}

// 获取当前登录的账户信息，需配合中间件使用
func CurrentAuth(c *gin.Context) Current {

	var accountID uint64
	if val, ok := c.Get("jwtx.accountID"); ok && val != nil {
		accountID, _ = val.(uint64)
	}

	return Current{
		AccountID:     accountID,
		LoginGroup:    c.GetString("jwtx.loginGroup"),
		LoginTerminal: c.GetString("jwtx.loginTerminal"),
		MakeTokenIP:   c.GetString("jwtx.makeTokenIP"),
	}
}
