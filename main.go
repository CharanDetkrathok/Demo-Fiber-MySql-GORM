package main

import (
	"demo-fiber-mysql-gorm/config"
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/rediscache"
	"demo-fiber-mysql-gorm/server"

	"github.com/gofiber/fiber/v2"
)

func init() {

	config.ConfigInit()
	database.InitDatabase()
	rediscache.InitRedisCache()

	// zaplogger.InitLogger()
	// newRedis := rediscache.MakeCache{
	// 	Key:    "KEY_zap_logger",
	// 	Data:   "DATA_zap_logger",
	// 	Expire: "1h",
	// }
	// zaplogger.ZapLog.Infof("Redis Struct %+v", newRedis)
	// zaplogger.ZapLog.Warnf("Redis Struct %+v", newRedis)
	// zaplogger.ZapLog.Errorf("Redis Struct %+v", newRedis)
}

func main() {

	defer rediscache.RedisCaching.RedisClient.Close()

	app := fiber.New()

	server := server.NewServer()
	server.RouterGroup(app)

	app.Listen(config.Env.API_PORT)

}
