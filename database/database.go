package database

import (
	"context"
	"demo-fiber-mysql-gorm/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

type SqlLogger struct {
	logger.Interface
}

func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", config.Env.MYSQL_USERNAME, config.Env.MYSQL_PASSWORD, config.Env.MYSQL_HOSTNAME, config.Env.MYSQL_DB_NAME)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: &SqlLogger{}},
	)
	if err != nil {
		panic(err)
	}

	Connection = db
}

func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, _ := fc()
	fmt.Printf("\n => SQL Statement :: %v \n\n", sqlStatement)
}
