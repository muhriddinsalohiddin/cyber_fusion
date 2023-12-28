package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateUser(c *fiber.Ctx) error {
	var u models.User
	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	fmt.Println(u)
	return handlerResponse(c, http.StatusCreated, u)
}

func (a *Api) GetUser(c *fiber.Ctx) error {
	limit := c.Query("limit")
	return handlerResponse(c, http.StatusOK, "GET: "+limit)
}

func (a *Api) GetByIdUser(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("idddd", id)
	return handlerResponse(c, http.StatusOK, "aka man yubordim:"+id)
}
func (a *Api) UpdateUser(c *fiber.Ctx) error {
	return handlerResponse(c, http.StatusOK, "Updated")
}

func (a *Api) DeleteUser(c *fiber.Ctx) error {
	return handlerResponse(c, http.StatusNoContent, nil)
}
