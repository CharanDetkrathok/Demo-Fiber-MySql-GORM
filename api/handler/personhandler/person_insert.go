package personhandler

import (
	"demo-fiber-mysql-gorm/model/request"
	"demo-fiber-mysql-gorm/model/response"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandler) InsertPerson_GORM(c *fiber.Ctx) error {

	newPerson := new(request.Person)
	if err := c.BodyParser(&newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "Insert GORM Bad request.",
		}
		return c.JSON(res)
	}

	if err := handle.PersonService.InsertPerson_GORM(newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM102",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "Insert GORM  Failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "GORM102",
		Data:         "Insert GORM Successfully.",
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(res)
}
