package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateAuthor(c *fiber.Ctx) error {
	var b models.Author
	err := c.BodyParser(&b)
	if err!=nil{
		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: " + err.Error())
	}
	fmt.Println(b)
	return handlerResponse(c, http.StatusCreated, "SUCCESS")
}


