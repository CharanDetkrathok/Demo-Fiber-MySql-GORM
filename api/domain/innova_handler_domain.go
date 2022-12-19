package domain

import "github.com/gofiber/fiber/v2"

type InnovaHandler interface {
	Get_InnovaOrder_With_InnovaOrderID(c *fiber.Ctx) error
}
