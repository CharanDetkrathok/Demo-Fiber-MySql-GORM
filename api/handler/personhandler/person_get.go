package personhandler

import (
	"demo-fiber-mysql-gorm/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandler) GetPersonWithPersonID_GORM(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	result, err := handle.PersonService.GetPersonWithPersonID_GORM(personId)
	if err != nil {
		res := response.Response{
			Code:         "-GROM101",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "Get GORM Failed.",
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
