package server

import (
	"demo-fiber-mysql-gorm/domain"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	personhandle "demo-fiber-mysql-gorm/handle/person"
	personrepo "demo-fiber-mysql-gorm/repository/person"
	personservice "demo-fiber-mysql-gorm/service/person"
)

type (
	server struct {
		PersonHandle domain.PersonHandle
	}

	Server interface {
		Router(app *fiber.App)
		personGroup(app *fiber.App)
	}
)

func NewServer(gorm_db *gorm.DB, redis *redis.Client) Server {

	newPersonRepo := personrepo.NewPersonRepo(gorm_db)
	newPersonService := personservice.NewPersonService(redis, newPersonRepo)
	newPersonHandle := personhandle.NewPersonHandle(newPersonService)

	return &server{newPersonHandle}
}
