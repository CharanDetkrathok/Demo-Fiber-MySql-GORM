package server

import "github.com/gofiber/fiber/v2"

func (server *server) personGroup(person fiber.Router) {

	person.Get("/:personId", server.PersonHandle.GetPersonWithPersonID_GORM)
	person.Post("", server.PersonHandle.InsertPerson_GORM)
	person.Put("", server.PersonHandle.UpdatePerson_GORM)
	person.Delete("/:personId", server.PersonHandle.DeletePerson_GORM)

}