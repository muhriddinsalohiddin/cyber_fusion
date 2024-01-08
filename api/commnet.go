package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateComment(c *fiber.Ctx) error {
	var u models.Comment

	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	err = a.stg.Comment.Create(&u)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err)
	}
	return handlerResponse(c, http.StatusCreated, u)

}
func (a *Api) GetCommentlist(c *fiber.Ctx) error {
	var req = models.Comment{
		UserId: c.Query("user_id"),
	}

	m, err := a.stg.Comment.Getlist(&req)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, m)
}

func (a *Api) UpdateComment(c *fiber.Ctx) error {
	var m models.Comment
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Updatedagi: "+err.Error())
	}
	id := c.Params("id")
	err = a.stg.Comment.Update(&m, &id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}

func (a *Api) DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.stg.Comment.DeleteComment(&id)

	if err != nil {
		fmt.Println("DeleteComment funksiyasida xato: ", err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}
