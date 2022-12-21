package innovahandler

import (
	"demo-fiber-mysql-gorm/model/response"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// get 1 condition json object
func (innovaHandler *innovaHandler) Get_InnovaOrder_With_InnovaOrderID(c *fiber.Ctx) error {

	innovaOrderID, err := strconv.Atoi(c.Params("innova_order_id"))
	if err != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         err.Error(),
			HttpCode:     201,
			ErrorMessage: "GET Innova Order BY GORM Bad request.",
		}
		return c.JSON(res)
	}

	result, err := innovaHandler.InnovaService.Get_InnovaOrder_With_InnovaOrderID(&innovaOrderID)
	if err != nil {
		res := response.Response{
			Code:         "-GORM102",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "GET Order BY GORM  Failed.",
		}
		return c.JSON(res)
	}

	resp := response.Response{
		Code:         "GORM101",
		Data:         result,
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(resp)
}

func (innovaHandler *innovaHandler) Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(c *fiber.Ctx) error {

	InnovaOrderID1, err1 := strconv.Atoi(c.Params("innova_order_id1"))
	InnovaOrderID2, err2 := strconv.Atoi(c.Params("innova_order_id2"))
	InnovaOrderID3, err3 := strconv.Atoi(c.Params("innova_order_id3"))
	if err1 != nil || err2 != nil || err3 != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         fmt.Sprintf("%v%v%v", err1.Error(), err2.Error(), err3.Error()),
			HttpCode:     201,
			ErrorMessage: "GET Innova Order BY GORM Bad request.",
		}
		return c.JSON(res)
	}

	result, err := innovaHandler.InnovaService.Get_InnovaOrder_With_InnovaOrderID_Slice_WhereIn(&InnovaOrderID1, &InnovaOrderID2, &InnovaOrderID3)
	if err != nil {
		res := response.Response{
			Code:         "-GORM102",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "GET Order BY GORM  Failed.",
		}
		return c.JSON(res)
	}

	resp := response.Response{
		Code:         "GORM101",
		Data:         result,
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(resp)
}

func (innovaHandler *innovaHandler) Get_InnovaOrder_With_InnovaOrderID_Slice(c *fiber.Ctx) error {

	system_brand_id, err := strconv.Atoi(c.Params("system_brand_id"))
	if err != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         err.Error(),
			HttpCode:     201,
			ErrorMessage: "GET Innova Order BY GORM Bad request.",
		}
		return c.JSON(res)
	}

	result, err := innovaHandler.InnovaService.Get_InnovaOrder_With_InnovaOrderID_Slice(&system_brand_id)
	if err != nil {
		res := response.Response{
			Code:         "-GORM102",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "GET Order BY GORM  Failed.",
		}
		return c.JSON(res)
	}

	resp := response.Response{
		Code:         "GORM101",
		Data:         result,
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(resp)
}

func (innovaHandler *innovaHandler) Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(c *fiber.Ctx) error {

	InnovaOrderID, err1 := strconv.Atoi(c.Params("innova_order_id"))
	SystemBrandID, err2 := strconv.Atoi(c.Params("system_brand_id"))
	AppMemberGroupID, err3 := strconv.Atoi(c.Params("app_member_group"))
	if err1 != nil || err2 != nil || err3 != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         fmt.Sprintf("%v%v%v", err1.Error(), err2.Error(), err3.Error()),
			HttpCode:     201,
			ErrorMessage: "GET Innova Order BY GORM Bad request.",
		}
		return c.JSON(res)
	}

	result, err := innovaHandler.InnovaService.Get_InnovaOrder_With_InnovaOrderID_SystemBrandID_AppMemberGroupID(&InnovaOrderID, &SystemBrandID, &AppMemberGroupID)
	if err != nil {
		res := response.Response{
			Code:         "-GORM102",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "Innova Order BY GORM  Failed.",
		}
		return c.JSON(res)
	}

	resp := response.Response{
		Code:         "GORM101",
		Data:         result,
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(resp)
}
