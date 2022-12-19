package server

import (
	"demo-fiber-mysql-gorm/domain"

	"github.com/gofiber/fiber/v2"

	personhandle "demo-fiber-mysql-gorm/handle/person"
	personrepo "demo-fiber-mysql-gorm/repository/person"
	personservice "demo-fiber-mysql-gorm/service/person"
)

type (
	server struct {
		PersonHandle domain.PersonHandle
	}

	Server interface {
		RouterGroup(app *fiber.App)
	}
)

func NewServer() Server {

	newPersonRepo := personrepo.NewPersonRepo()
	newPersonService := personservice.NewPersonService(newPersonRepo)
	newPersonHandle := personhandle.NewPersonHandle(newPersonService)

	return &server{newPersonHandle}
}
