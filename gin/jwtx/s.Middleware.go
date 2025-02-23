package jwtx

import (
	"context"
	"errors"
	"time"

	"github.com/5-say/go-tool/gin/jwtx/db/dao/model"
	"github.com/5-say/go-tool/gin/jwtx/db/dao/query"
	"github.com/5-say/go-tool/gin/jwtx/tool"
	"github.com/5-say/go-tool/logx"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 中间件认证失败后的处理函数 ..
var MiddlewareFailHandlerFunc = func(c *gin.Context, err error) {
	logx.Debug().Err(err)
	c.AbortWithStatus(401)
}

// jwtx token 中间件（token 解析、校验、刷新） ..
//
//	requestGroup string 请求访问的分组
//
// e.g.
//
//	jwtx.Singleton.Middleware("admin")
//	jwtx.Singleton.Middleware("merchant")
func (s *SingletonT) Middleware(requestGroup string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 解析 token
		claims, err := tool.ParseToken(c.GetHeader("Authorization"), &s.PrivateKey[requestGroup].PublicKey)
		if err != nil {
			MiddlewareFailHandlerFunc(c, err)
			return
		}

		// 提取 claims
		var (
			now    = time.Now()                      // 当前时间
			iat, _ = claims.GetIssuedAt()            // 签发时间
			exp, _ = claims.GetExpirationTime()      // 过期时间
			tid    = uint64(claims["tid"].(float64)) // token ID
		)

		// token 过期时间校验
		if exp.Unix() < now.Unix() {
			MiddlewareFailHandlerFunc(c, errors.New("token has expired"))
			return
		}

		// 数据库中查找 token
		token, err := s.getDBToken(requestGroup, tid)
		if err != nil {
			MiddlewareFailHandlerFunc(c, err)
			return
		}

		// 校验数据库 token 信息
		err = s.checkDBToken(requestGroup, token, now, c.ClientIP())
		if err != nil {
			MiddlewareFailHandlerFunc(c, err)
			return
		}

		// 刷新 token
		newToken, err := s.refreshToken(requestGroup, token, now, iat)
		if err != nil {
			MiddlewareFailHandlerFunc(c, err)
			return
		}

		// 将当前登录的账户信息存储上下文
		c.Set("jwtx.accountID", token.AccountID)         // 账户 ID
		c.Set("jwtx.loginGroup", token.LoginGroup)       // 登录的分组
		c.Set("jwtx.loginTerminal", token.LoginTerminal) // 登录的终端
		c.Set("jwtx.makeTokenIP", token.MakeTokenIP)     // 首次请求生成 token 的 IP 地址

		// 请求前
		c.Next()
		// 请求后

		// 响应头填充刷新后的 token
		c.Writer.Header().Add("token", newToken)
	}
}

// 获取数据库 token 信息
func (s *SingletonT) getDBToken(group string, tid uint64) (token *model.JwtxToken, err error) {
	// 查找 token
	q := query.Use(s.DB[group])
	o := q.JwtxToken
	token, err = o.WithContext(context.Background()).Where(o.ID.Eq(tid)).First()
	if err != nil {
		return nil, err
	}
	return token, nil
}

// 校验数据库 token 信息
func (s *SingletonT) checkDBToken(group string, token *model.JwtxToken, now time.Time, clientIP string) (err error) {
	// 分组校验
	if token.LoginGroup != group {
		return errors.New("auth group fail")
	}
	// 过期时间校验
	if token.ExpirationAt.Unix() < now.Unix() {
		return errors.New("the token has expired")
	}
	// IP 一致性校验
	if s.Config[group].CheckIP {
		if clientIP != token.MakeTokenIP {
			return errors.New("client ip is changed, please login again")
		}
	}
	return
}

// 自动刷新 token
func (s *SingletonT) refreshToken(group string, token *model.JwtxToken, now time.Time, iat *jwt.NumericDate) (newToken string, err error) {

	var config = s.Config[group]

	if iat.Unix() == token.FinalRefreshAt.Unix() { // token 未刷新

		// 原始的 token 过期时间
		expTime := token.ExpirationAt

		// 需要刷新 token
		if iat.Unix()+config.RefreshInterval < now.Unix() {

			// 自动续期
			if config.AutomaticRenewal {
				expTime = now.Add(time.Duration(config.AccessExpireByHour) * time.Hour)
			}

			// 构造 token 字符串（过期时间不变，签发时间顺延）
			newToken, err = tool.GenerateToken(s.PrivateKey[group], jwt.MapClaims{
				"iat": now.Unix(),     // 签发时间
				"exp": expTime.Unix(), // 过期时间
				"tid": token.ID,       // jwt token ID
			}, tool.P_384)
			if err != nil {
				return "", err
			}

			// 更新数据库
			q := query.Use(s.DB[group])
			o := q.JwtxToken
			_, err = o.WithContext(context.Background()).Where(o.ID.Eq(token.ID)).UpdateSimple(
				o.LastRefreshAt.Value(token.FinalRefreshAt),
				o.FinalRefreshAt.Value(now),
			)
			if err != nil {
				return "", err
			}
		}

	} else if iat.Unix() == token.LastRefreshAt.Unix() { // token 已刷新

		// 当前时间 超出 并发容错时间（不允许继续使用）
		if now.Unix() > token.FinalRefreshAt.Unix()+config.FaultTolerance {
			return "", errors.New("out of concurrent fault tolerance time")
		}
	}

	return
}
