package server

import "github.com/gofiber/fiber/v2"

func (server *server) innovaGroup(innova fiber.Router) {

	innova.Get("/:innova_order_id", server.InnovaHandler.Get_InnovaOrder_With_InnovaOrderID)

}
