package api

import (
	"app/models"
	//"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateNotification(c *fiber.Ctx) error {
	var u models.Notification

	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	a.stg.Notification.Create(&u)
	return handlerResponse(c, http.StatusCreated, u)

}
func (a *Api) GetnotificationList(c *fiber.Ctx) error {
	var m models.List
	err := a.stg.List.Getlist(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c,http.StatusOK, m)
}


func (a *Api) UpdateNotification(c *fiber.Ctx) error {
	var m models.Notification
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Updatedagi: "+err.Error())
	}
	err = a.stg.Notification.Update(&m)
	if err!=nil{
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}

func (a *Api) DeleteNotification(c *fiber.Ctx) error {
	var m models.Notification
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Deletedagi: "+err.Error())
	}
	err = a.stg.Notification.Delete(&m)
	if err!=nil{
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}
