package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) File(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return handlerResponse(c, fiber.StatusBadRequest, err.Error())
	}

	destination := fmt.Sprintf("./images/%s", file.Filename)
	err = c.SaveFile(file, destination)
	if err != nil {
		return handlerResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return handlerResponse(c, fiber.StatusOK, "OK")
}
