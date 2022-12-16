package main

import (
	"demo-fiber-mysql-gorm/config"
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/server"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var gorm_db *gorm.DB
var redis_cache *redis.Client
var port string

func init() {

	environments := config.ConfigInit()
	newDatabase := database.NewDatabase(&environments)
	gorm_db = newDatabase.GormInit()
	redis_cache = newDatabase.RedisInit()
	port = environments.API_PORT
}

func main() {

	defer redis_cache.Close()

	app := fiber.New()

	server := server.NewServer(gorm_db, redis_cache)
	server.Router(app)

	app.Listen(port)

}
