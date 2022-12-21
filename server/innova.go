package server

import "github.com/gofiber/fiber/v2"

func (server *server) innovaGroup(innova fiber.Router) {

	innova.Get("/:innova_order_id", server.InnovaHandler.Get_InnovaOrder_With_InnovaOrderID)
	innova.Get("/slice-where-in/:innova_order_id1/:innova_order_id2/:innova_order_id3", server.InnovaHandler.Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn)
	innova.Get("/slice/:system_brand_id", server.InnovaHandler.Get_InnovaOrder_With_InnovaOrderID_Slice)
	innova.Get("/slice/:innova_order_id/:system_brand_id/:app_member_group", server.InnovaHandler.Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID)

}
