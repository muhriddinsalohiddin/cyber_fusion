package api

import (
	"app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (a *Api) CreateUser(c *fiber.Ctx) error {
	var u models.User
	err := c.BodyParser(&u)
	u.Id = uuid.NewString()
	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	a.stg.Notification.Create(&u)
	return handlerResponse(c, http.StatusCreated, u)
}

func (a *Api) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := a.stg.User.Get(id)
	return handlerResponse(c, http.StatusCreated, user)
}

func (a *Api) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := a.stg.User.Get(id)
	return handlerResponse(c, http.StatusCreated, user)
}
func (a *Api) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := a.stg.User.Get(id)
	return handlerResponse(c, http.StatusCreated, user)
}
