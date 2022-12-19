package server

import (
	"demo-fiber-mysql-gorm/api/domain"
	"demo-fiber-mysql-gorm/api/handler/innovahandler"
	"demo-fiber-mysql-gorm/api/handler/personhandler"
	"demo-fiber-mysql-gorm/api/repository/innovarepo"
	"demo-fiber-mysql-gorm/api/repository/personrepo"
	"demo-fiber-mysql-gorm/api/service/innovaservice"
	"demo-fiber-mysql-gorm/api/service/personservice"

	"github.com/gofiber/fiber/v2"
)

type (
	server struct {
		PersonHandle  domain.PersonHandler
		InnovaHandler domain.InnovaHandler
	}

	Server interface {
		RouterGroup(app *fiber.App)
	}
)

func NewServer() Server {

	newPersonRepo := personrepo.NewPersonRepo()
	newPersonService := personservice.NewPersonService(newPersonRepo)
	newPersonHandle := personhandler.NewPersonHandler(newPersonService)

	newInnovaRepo := innovarepo.NewInnovaRepo()
	newInnovaService := innovaservice.NewInnovaService(newInnovaRepo)
	newInnovaHandler := innovahandler.NewInnovaHandler(newInnovaService)

	return &server{
		newPersonHandle,
		newInnovaHandler,
	}
}
