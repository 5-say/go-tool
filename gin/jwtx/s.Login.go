package jwtx

import (
	"errors"
	"time"

	"github.com/5-say/go-tool/gin/jwtx/db/dao/model"
	"github.com/5-say/go-tool/gin/jwtx/db/dao/query"
	"github.com/5-say/go-tool/gin/jwtx/tool"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// jwt token 登录
//
//	c             *gin.Context
//	loginGroup    string       登录的 token 分组
//	loginTerminal string       登录的 终端名称（后续登出时，若引用此处的传值为“all”，则无视配置，强制登出所有终端）
//	accountID     uint64       登录的 账户 ID
//
// e.g.
//
//	jwtx.Singleton.Login(c, "admin", "pc", 1)
//	jwtx.Singleton.Login(c, "admin", "mobile", 1)
func (s *SingletonT) Login(c *gin.Context, loginGroup string, loginTerminal string, accountID uint64) (err error) {
	// 获取分组配置
	var (
		config     = s.Config[loginGroup]
		privateKey = s.PrivateKey[loginGroup]
		q          = query.Use(s.DB[loginGroup])
	)

	// 清除旧 token
	err = s.Logout(c, loginGroup, loginTerminal, accountID)
	if err != nil {
		return
	}

	var (
		now     = time.Now()
		expTime = now.Add(time.Hour * time.Duration(config.AccessExpireByHour))
	)

	// 创建新 token
	token := model.JwtxToken{
		AccountID:      accountID,
		LoginGroup:     loginGroup,
		LoginTerminal:  loginTerminal,
		MakeTokenIP:    c.ClientIP(),
		CreatedAt:      now,
		LastRefreshAt:  now,
		FinalRefreshAt: now,
		ExpirationAt:   expTime,
	}
	if q.JwtxToken.WithContext(c).Create(&token) != nil {
		return errors.New("create token fail")
	}

	// 构造 token 字符串
	tokenStr, err := tool.GenerateToken(privateKey, jwt.MapClaims{
		"iat": now.Unix(),     // 签发时间
		"exp": expTime.Unix(), // 过期时间
		"tid": token.ID,       // jwt token ID
	}, tool.P_384)
	if err != nil {
		return errors.New("token refresh fail")
	}

	// 响应头填充 token
	c.Writer.Header().Add("token", tokenStr)

	return
}
