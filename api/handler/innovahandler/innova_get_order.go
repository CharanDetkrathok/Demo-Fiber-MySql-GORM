package innovahandler

import (
	"demo-fiber-mysql-gorm/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
