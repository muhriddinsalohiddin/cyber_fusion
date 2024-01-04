package api

import (
	"app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateMessage(c *fiber.Ctx) error {
	var m models.Message
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	err = a.stg.Message.Create(m)
	if err!=nil{
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusCreated, "SUCCESS")
}
func (a *Api) UpdateMessage(c *fiber.Ctx) error {
	var m models.Message
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik updatedagi: "+err.Error())
	}
	err = a.stg.Message.Update(m)
	if err!=nil{
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}
