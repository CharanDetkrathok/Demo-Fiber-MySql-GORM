package database

import (
	"context"
	"demo-fiber-mysql-gorm/config"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	connection struct {
		Env *config.Config
	}

	Connection interface {
		GormInit() *gorm.DB
		RedisInit() *redis.Client
	}

	SqlLogger struct {
		logger.Interface
	}
)

func NewDatabase(environments *config.Config) Connection {
	return &connection{environments}
}

func (conn *connection) GormInit() *gorm.DB {

	// * ------------ ใช้ gorm.Dialector ในการเชื่อมต่อกับ GORM โดยใช้ package "gorm.io/driver/mysql" ---------------
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", conn.Env.MYSQL_USERNAME, conn.Env.MYSQL_PASSWORD, conn.Env.MYSQL_HOSTNAME, conn.Env.MYSQL_DB_NAME)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: &SqlLogger{}},
	)
	if err != nil {
		panic(err)
	}

	return db
}

func (conn *connection) RedisInit() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conn.Env.REDIS_ADDRESS,
		Password: conn.Env.REDIS_PASSWORD,
		DB:       conn.Env.REDIS_DB_NUM,
	})
}

func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, _ := fc()
	fmt.Printf("\n => SQL Statement :: %v \n\n", sqlStatement)
}
