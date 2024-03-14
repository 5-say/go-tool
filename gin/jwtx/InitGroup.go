package jwtx

import (
	"crypto/ecdsa"
	"time"

	"github.com/5-say/go-tool/gin/jwtx/tool"
	"github.com/5-say/go-tool/gorm/mysqlx"
	"github.com/5-say/go-tool/logx"
	"github.com/jinzhu/configor"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化分组（支持多次调用）
//
//	group          string 分组名称
//	configPath     string 分组配置文件路径
//	privateKeyPath string 分组私钥文件路径
//
// e.g.
//
//	jwtx.InitGroup("admin", "jwtx/config/admin.yaml", "jwtx/config/admin.key")
//	jwtx.InitGroup("user",  "jwtx/config/user.yaml",  "jwtx/config/user.key")
func InitGroup(group, configPath, privateKeyPath string) *SingletonMode {
	// 取得单例
	if Singleton == nil {
		Singleton = &SingletonMode{}
		Singleton.DB = make(map[string]*gorm.DB)
		Singleton.Config = make(map[string]GroupConfig)
		Singleton.PrivateKey = make(map[string]*ecdsa.PrivateKey)
	}

	// 引用简化
	var s = Singleton

	// 初始化分组配置
	var config GroupConfig
	err := configor.Load(&config, configPath)
	if err != nil {
		panic(err)
	}
	s.Config[group] = config

	// 初始化分组数据库连接
	s.DB[group] = mysqlx.New(config.MysqlDSN, logx.GormLogger(logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		LogLevel:                  logger.Warn,
	}))

	// 初始化分组私钥
	privateKey, err := tool.GetPrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	s.PrivateKey[group] = privateKey

	return s
}
