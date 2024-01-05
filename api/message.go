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
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Create: "+err.Error())
	}
	err = a.stg.Message.Create(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusCreated, "SUCCESS")
}

func (a *Api) UpdateMessage(c *fiber.Ctx) error {
	var m models.Message
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Updatedagi: "+err.Error())
	}
	err = a.stg.Message.Update(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}

func (a *Api) DeleteMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.stg.Message.Delete(&id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}

func (a *Api) GetMessageList(c *fiber.Ctx) error {
	var m models.List
	err := a.stg.List.GetMessageList(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, m)
}

func (a *Api) GetMessageById(c *fiber.Ctx) error {
	var m models.Message
	id := c.Params("id")
	err := a.stg.Message.GetMessage(&m, &id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusAccepted,m)
}
