package server

import (
	"github.com/gofiber/fiber/v2"
)

func (server *server) RouterGroup(app *fiber.App) {

	personGormGroup := app.Group("/person_gorm")
	server.personGroup(personGormGroup)

	innovaGormGroup := app.Group("/innova_gorm")
	server.innovaGroup(innovaGormGroup)

}
