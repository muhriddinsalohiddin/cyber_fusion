package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateAuthor(c *fiber.Ctx) error {
	var u models.Author
	err := c.BodyParser(&u)
	if err!=nil{
		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: " + err.Error())
	}
	fmt.Println(u)
	return handlerResponse(c, http.StatusCreated, u)
}


