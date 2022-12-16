package domain

import "github.com/gofiber/fiber/v2"

type PersonHandle interface {
	GetPersonWithPersonID_GORM(c *fiber.Ctx) error
	InsertPerson_GORM(c *fiber.Ctx) error
	UpdatePerson_GORM(c *fiber.Ctx) error
	DeletePerson_GORM(c *fiber.Ctx) error
}
