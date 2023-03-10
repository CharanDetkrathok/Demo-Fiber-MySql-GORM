package personhandler

import (
	"demo-fiber-mysql-gorm/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandler) DeletePerson_GORM(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	if err := handle.PersonService.DeletePerson_GORM(personId); err != nil {
		res := response.Response{
			Code:         "-GORM104",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "Delete GORM  Failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "GORM104",
		Data:         "Delete GORM Successfully.",
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(res)
}
