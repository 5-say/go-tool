package jwtx

import (
	"errors"

	"github.com/5-say/go-tool/gin/jwtx/db/dao/query"
	"github.com/5-say/go-tool/logx"
	"github.com/gin-gonic/gin"
)

// jwt token 登出
//
//	c             *gin.Context
//	loginGroup    string       登录的 token 分组
//	loginTerminal string       登录的 终端名称（当传值为“all”时，无视配置，强制登出所有终端）
//	accountID     uint32       登录的 账户 ID
//
// e.g.
//
//	jwtx.Singleton.Logout(c, "admin", "pc", 1)
//	jwtx.Singleton.Logout(c, "admin", "all", 1)
func (s *SingletonMode) Logout(c *gin.Context, loginGroup string, loginTerminal string, accountID uint32) (err error) {
	// 获取分组配置
	var (
		q      = query.Use(s.DB[loginGroup])
		config = s.Config[loginGroup]
	)

	if config.SingleEnd || loginTerminal == "all" { // 强制单端登录

		if _, err := q.JwtxToken.WithContext(c).Where(
			q.JwtxToken.AccountID.Eq(accountID),
			q.JwtxToken.LoginGroup.Eq(loginGroup),
		).Delete(); err != nil {
			logx.Error().Msg("强制单端登录，清除旧 token 失败 [sqlfail: " + err.Error() + "]")
			return errors.New("clear token fail")
		}

	} else { // 多端登录

		if _, err := q.JwtxToken.WithContext(c).Where(
			q.JwtxToken.AccountID.Eq(accountID),
			q.JwtxToken.LoginGroup.Eq(loginGroup),
			q.JwtxToken.LoginTerminal.Eq(loginTerminal),
		).Delete(); err != nil {
			logx.Error().Msg("多端登录，清除旧 token 失败 [sqlfail: " + err.Error() + "]")
			return errors.New("clear token fail")
		}
	}

	return
}
