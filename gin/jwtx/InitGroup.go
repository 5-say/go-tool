package jwtx

import (
	"crypto/ecdsa"
	"time"

	"github.com/5-say/go-tool/crypto/ecdsatool"
	"github.com/5-say/go-tool/gin/jwtx/db/dao/model"
	"github.com/5-say/go-tool/gin/jwtx/tool"
	"github.com/5-say/go-tool/gorm/mysqlx"
	"github.com/5-say/go-tool/logx"
	"github.com/5-say/go-tool/utilx"
	"github.com/jinzhu/configor"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "embed"
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
func InitGroup(group, configPath, privateKeyPath string) *SingletonT {
	// 配置文件不存在则创建
	createConfigFileIfNotExist(configPath)

	// 私钥文件不存在则创建
	createKeyFileIfNotExist(privateKeyPath)

	// 取得单例
	if Singleton == nil {
		Singleton = &SingletonT{}
		Singleton.DB = make(map[string]*gorm.DB)
		Singleton.Config = make(map[string]GroupConfig)
		Singleton.PrivateKey = make(map[string]*ecdsa.PrivateKey)
	}

	// 引用简化
	var s = Singleton

	// 初始化分组配置，每秒自动更新
	var config GroupConfig
	err := configor.New(&configor.Config{AutoReload: true}).Load(&config, configPath)
	if err != nil {
		panic(err)
	}
	s.Config[group] = config

	// 初始化分组数据库连接
	s.DB[group], err = mysqlx.New(config.MysqlDSN, logx.GormLogger(logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		LogLevel:                  logger.Warn,
	}))
	if err != nil {
		panic(err)
	}

	// 初始化数据库结构
	s.DB[group].AutoMigrate(&model.JwtxToken{})

	// 初始化分组私钥
	privateKey, err := tool.GetPrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	s.PrivateKey[group] = privateKey

	return s
}

//go:embed example/demo.yaml
var demoConfig string

func createConfigFileIfNotExist(filePath string) {
	utilx.ForceCreateWriteIfNotExist(filePath, 0755, demoConfig)
}

func createKeyFileIfNotExist(filePath string) {
	key, _ := ecdsatool.GenerateKey(ecdsatool.P_384)
	privateKeyPEMBytes, _ := ecdsatool.PEM_EncodePrivateKey(key)
	content := "\n"
	content += "密钥格式: NIST P-384"
	content += "\n\n"
	content += string(privateKeyPEMBytes)
	content += "\n"
	content += "密钥格式: NIST P-384"
	content += "\n"

	utilx.ForceCreateWriteIfNotExist(filePath, 0755, content)
}
