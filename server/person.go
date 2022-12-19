package server

import "github.com/gofiber/fiber/v2"

func (server *server) personGroup(app fiber.Router) {

	app.Get("/:personId", server.PersonHandle.GetPersonWithPersonID_GORM)
	app.Post("", server.PersonHandle.InsertPerson_GORM)
	app.Put("", server.PersonHandle.UpdatePerson_GORM)
	app.Delete("/:personId", server.PersonHandle.DeletePerson_GORM)

}