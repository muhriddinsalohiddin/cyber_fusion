package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreatePost(c *fiber.Ctx) error {
	var m models.Post
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik Create: "+err.Error())
	}
	fmt.Println("post", m)
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

	err := a.stg.Post.Delete(&models.Post{Id: c.Params("id")})
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}
func (a *Api) GetPost(c *fiber.Ctx) error {
	m, err := a.stg.Post.GetPostlist()
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, m)
}
func (a *Api) GetByIdPost(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := a.stg.Post.GetByIdPost(id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, post)
}
