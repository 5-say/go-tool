package mysqlx

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 建立新的数据库连接
//
//	dsn       string
//	useLogger logger.Interface
//
// e.g.
//
//	mysqlx.New("", logger.Discard)
//	mysqlx.New("", logger.Default.LogMode(logger.Error))
//	mysqlx.New("", logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{}))
func New(dsn string, useLogger logger.Interface) *gorm.DB {
	// var dsn = config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 是否跳过 “根据当前 MySQL 版本自动配置”
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,      // 迁移时禁用外键约束
		SkipDefaultTransaction:                   true,      // 跳过默认事务
		Logger:                                   useLogger, // 日志
	})

	if err != nil {
		log.Fatalf("Fail to db logger: %v\n", err)
	}

	return db
}
