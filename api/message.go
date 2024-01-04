package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateMessage(c *fiber.Ctx) error {
	var m models.Message
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	fmt.Println(m)
	return handlerResponse(c, http.StatusCreated, m)
}
