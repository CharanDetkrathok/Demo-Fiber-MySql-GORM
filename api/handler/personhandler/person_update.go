package personhandler

import (
	"demo-fiber-mysql-gorm/model/request"
	"demo-fiber-mysql-gorm/model/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandler) UpdatePerson_GORM(c *fiber.Ctx) error {

	newPerson := new(request.Person) // แบบที่ 1
	// var newPerson *request.Person // แบบที่ 2
	if err := c.BodyParser(&newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         err,
			HttpCode:     200,
			ErrorMessage: "Update GORM Bad request.",
		}
		return c.JSON(res)
	}

	fmt.Println(newPerson.Marshal())

	if err := handle.PersonService.UpdatePerson_GORM(newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM103",
			Data:         err.Error(),
			HttpCode:     200,
			ErrorMessage: "Update GORM  Failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "GORM103",
		Data:         "Update GORM Successfully.",
		HttpCode:     200,
		ErrorMessage: "",
	}

	return c.JSON(res)
}
