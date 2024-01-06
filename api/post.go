package api

import (
	"app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreatePost(c *fiber.Ctx) error {
	var m models.Post
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Create: "+err.Error())
	}
	err = a.stg.Post.Create(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusCreated, "SUCCESS CREATED")
}

func (a *Api) UpdatePost(c *fiber.Ctx) error {
	var m models.Post
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Updatedagi: "+err.Error())
	}
	err = a.stg.Post.Update(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}

func (a *Api) DeletePost(c *fiber.Ctx) error {
	var m models.Post
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Deletedagi: "+err.Error())
	}
	err = a.stg.Post.Delete(&m)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}
func (a *Api) GetPost(c *fiber.Ctx) error {
	m, err := a.stg.List.GetPostlist()
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, m)
}
func (a *Api) GetByIdPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := a.stg.List.GetByIdPost(id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, post)
}
