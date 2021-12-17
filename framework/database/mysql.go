package database

import (
	"os"

	"go.uber.org/zap"
	"goku.net/framework/commons"
	"goku.net/framework/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDBs map[string]*gorm.DB
)

func Init(configs []*config.MysqlConfig) {
	mysqlDBs = make(map[string]*gorm.DB)
	for _, config := range configs {
		gormConfig := mysql.Config{
			DSN:                       config.Dsn(), // DSN data source name
			DefaultStringSize:         191,          // string 类型字段的默认长度
			DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,        // 根据版本自动配置
		}
		if db, err := gorm.Open(mysql.New(gormConfig), gormOptions()); err == nil {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(config.MaxIdleConns)
			sqlDB.SetMaxOpenConns(config.MaxOpenConns)
			mysqlDBs[config.DbName] = db
		} else {
			commons.Logger().Error("MySQL连接异常", zap.Any("err", err))
			os.Exit(-1)
			return
		}
	}
}

func gormOptions() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}

	return config
}

func GetDB(dbName string) *gorm.DB {
	if db, ok := mysqlDBs[dbName]; ok {
		return db
	}
	return nil
}
