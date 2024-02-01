package jwtx

import (
	"crypto/ecdsa"

	"github.com/5-say/go-tool/gin/jwtx/tool"
	"github.com/jinzhu/configor"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化分组（支持多次调用）
//
//	group          string 分组名称
//	configPath     string 分组配置文件路径
//	privateKeyPath string 分组私钥文件路径
//
// ex:
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
	gormdb, err := gorm.Open(mysql.Open(config.MysqlDSN))
	if err != nil {
		panic(err)
	}
	s.DB[group] = gormdb

	// 初始化分组私钥
	privateKey, err := tool.GetPrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	s.PrivateKey[group] = privateKey

	return s
}
