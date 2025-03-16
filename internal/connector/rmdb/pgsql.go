package rmdb

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var PgDB *gorm.DB

// InitDatabasePostgres initializes PostgreSQL database connection
func InitDatabasePostgres() (*gorm.DB, error) {
	if PgDB != nil {
		return PgDB, nil
	}

	dsn := "host=" + viper.GetString("Postgres.Host") +
		" user=" + viper.GetString("Postgres.User") +
		" password=" + viper.GetString("Postgres.Password") +
		" dbname=" + viper.GetString("Postgres.DbName") +
		" port=" + viper.GetString("Postgres.Port") +
		" sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("Postgres.TablePrefix"),
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	PgDB = db
	log.Println("PostgreSQL connected")
	return db, nil
}
