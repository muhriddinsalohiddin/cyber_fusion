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
	err=a.stg.Notification.Create(&u)
	if err!=nil{
		return handlerResponse(c,http.StatusInternalServerError,err)
	}
	return handlerResponse(c, http.StatusCreated, u)

}
func (a *Api) GetnotificationList(c *fiber.Ctx) error {
	
	m,err := a.stg.Notification.Getlist()
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
	id := c.Params("id")
	err = a.stg.Notification.Update(&id,&m)
	if err!=nil{
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}

// func (a *Api) DeleteNotification(c *fiber.Ctx) error {
// 	Id:=c.Params("id")
// 	err := a.stg.Notification.Delete(Id)
// 	if err!=nil{
// 		return handlerResponse(c, http.StatusInternalServerError, err.Error())
// 	}

// 	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
// }
func (a *Api) DeleteNotification(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.stg.Notification.Delete(&id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}