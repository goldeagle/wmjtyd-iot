package rmdb

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// connectDb
func InitDatabaseMysql() (*gorm.DB, error) {

	if DB != nil {
		return DB, nil
	}

	dsn := viper.GetString("Database.User") + ":" + viper.GetString("Database.Password") + "@tcp(" + viper.GetString("Database.Host") + ":" + viper.GetString("Database.Port") + ")/" +
		viper.GetString("Database.DbName") + "?charset=" + viper.GetString("Database.Chartset") + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    false, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   false, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}
	tbl_prefix := viper.GetString(("database.table_prefix"))
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tbl_prefix, // 表前缀
			SingularTable: true,       // 禁用表名复数
		},
		DisableForeignKeyConstraintWhenMigrating: true,                                // 禁用自动创建外键约束
		SkipDefaultTransaction:                   true,                                // 禁用默认事务，将获得大约 30%+ 性能提升
		PrepareStmt:                              true,                                // 创建并缓存预编译语句，可以提高后续的调用速度
		Logger:                                   logger.Default.LogMode(logger.Info), // 使用自定义 Logger
	})

	if err != nil {
		panic("Mysql 数据库链接失败gorm")
	}

	DB = db

	// 初始化数据表
	log.Println("Mysql connected")

	return db, nil
}
