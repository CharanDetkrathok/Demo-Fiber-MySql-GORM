package server

import "github.com/gofiber/fiber/v2"

func (server *server) personGroup(app *fiber.App) {

	app.Get("/person_gorm/:personId", server.PersonHandle.GetPersonWithPersonID_GORM)
	app.Post("/person_gorm", server.PersonHandle.InsertPerson_GORM)
	app.Put("/person_gorm", server.PersonHandle.UpdatePerson_GORM)
	app.Delete("/person_gorm/:personId", server.PersonHandle.DeletePerson_GORM)

}
