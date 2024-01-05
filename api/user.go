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
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 15: "+err.Error())
	}
	err = a.stg.User.Create(&u)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusCreated, "SUCCESS")
}

func (a *Api) GetUser(c *fiber.Ctx) error {
	fmt.Println("queries", c.Query("farruxjon"))
	fmt.Println("gender", c.Query("gender"))
	var u models.User
	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 28: "+err.Error())
	}

	user, err := a.stg.User.GetList(u.Id)
	if err != nil {
		return err
	}
	// a.stg.User.GetList("")
	return handlerResponse(c, http.StatusCreated, user)

	// return

}

func (a *Api) GetByIdUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := a.stg.User.GetList(id)
	if err != nil {
		return err
	}
	// a.stg.User.GetList("")
	return handlerResponse(c, http.StatusCreated, user)

}
func (a *Api) UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")
	user, err := a.stg.User.GetList(id)
	if err != nil {
		return err
	}
	return handlerResponse(c, http.StatusCreated, user)
}
func (a *Api) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, _ := a.stg.User.GetList(id)
	return handlerResponse(c, http.StatusCreated, user)
}
