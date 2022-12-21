package main

import (
	"demo-fiber-mysql-gorm/config"
	"demo-fiber-mysql-gorm/database"
	"demo-fiber-mysql-gorm/rediscache"
	"demo-fiber-mysql-gorm/server"

	"github.com/gofiber/fiber/v2"
)

//todo::===================== API DEMO IS USED BY GORM ========================
//*   :: done ==> get ด้วย 1 เงื่อนไข                	  ==> (โดยใช้ entity struct)
//*   :: wait ==> get ด้วย 3 เงื่อนไข                	  ==> (โดยใช้ entity struct)
//*   :: wait ==> get แบบเป็น slice     	 	   		==> (โดยใช้ entity struct)
//*   :: wait ==> get แบบเป็น slice where in + 3 เงื่อนไข ==> (โดยใช้ entity struct)
//?   :: wait ==> get แบบใช้ scope                 		==> (โดยใช้ entity struct)
//?   :: wait ==> get แบบใช้ scope      + 3 เงื่อนไข 	  ==> (โดยใช้ entity struct)
//?   :: wait ==> get แบบใช้ payload                    ==> (โดยใช้ entity struct)
//?   :: wait ==> get แบบใช้ payload    + 3 เงื่อนไข      ==> (โดยใช้ entity struct)
//?   :: wait ==> get แบบใช้ pagination                 ==> (โดยใช้ entity struct)
//?   :: wait ==> get แบบใช้ pagination + 3 เงื่อนไข      ==> (โดยใช้ entity struct)

//?   :: wait ==> insert
//?   :: wait ==> update
//?   :: wait ==> delete

func init() {
	config.ConfigInit()
	database.InitDatabase()
	rediscache.InitRedisCache()
}

func main() {
	defer rediscache.RedisCaching.RedisClient.Close()
	app := fiber.New()
	server := server.NewServer()
	server.RouterGroup(app)
	app.Listen(config.Env.API_PORT)
}
