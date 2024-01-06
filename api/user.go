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

func (a *Api) GetByIdUser(c *fiber.Ctx) error {
	var u models.User
	id := c.Params("id")
	fmt.Println(&u)
	// if err != nil {
	// 	return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 28: "+err.Error())
	// }
	user, err := a.stg.User.GetById(id)
	if err != nil {
		return err
	}
	// a.stg.User.GetById("")
	return handlerResponse(c, http.StatusCreated, user)

	// return

}
func (a *Api) GetUser(c *fiber.Ctx) error {
	user, err := a.stg.User.Get()
	if err != nil {
		return err
	}

	return handlerResponse(c, http.StatusCreated, user)

	// return

}
func (a *Api) UpdateUser(c *fiber.Ctx) error {
	var u models.User
	err:=c.BodyParser(&u)
	if err != nil {	
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 28: "+err.Error())
	}

	id := c.Params("id")
	 a.stg.User.Update(u,id)
	
	// return handlerResponse(c, http.StatusCreated, user)
	return handlerResponse(c, http.StatusCreated, "SUCCESS")
}
func (a *Api) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	 err := a.stg.User.Delete(id)
	 if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	 }

	return handlerResponse(c, http.StatusCreated, "TEG TUGI BILAN O'CHIP KETTI :)")
}
