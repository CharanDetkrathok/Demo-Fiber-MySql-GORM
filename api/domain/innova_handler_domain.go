package domain

import "github.com/gofiber/fiber/v2"

type InnovaHandler interface {
	Get_InnovaOrder_With_InnovaOrderID(c *fiber.Ctx) error
	Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(c *fiber.Ctx) error
	Get_InnovaOrder_With_InnovaOrderID_Slice(c *fiber.Ctx) error
	Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(c *fiber.Ctx) error
}
