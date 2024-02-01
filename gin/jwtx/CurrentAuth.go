package jwtx

import "github.com/gin-gonic/gin"

type Current struct {
	AccountID     uint32
	LoginGroup    string
	LoginTerminal string
}

// 获取当前登录的账户信息，需配合中间件使用
func CurrentAuth(c *gin.Context) Current {

	var accountID uint32
	if val, ok := c.Get("jwtx.accountID"); ok && val != nil {
		accountID, _ = val.(uint32)
	}

	return Current{
		AccountID:     accountID,
		LoginGroup:    c.GetString("jwtx.loginGroup"),
		LoginTerminal: c.GetString("jwtx.loginTerminal"),
	}
}
