package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateAuthor(c *fiber.Ctx) error {
	var b models.Author
	err := c.BodyParser(&b)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: "+err.Error())
	}
	fmt.Println(b)
	a.stg.Author.CreateAuthor(&b)
	return handlerResponse(c, http.StatusCreated, b)
}

func (a *Api) UpdateAuthor(c *fiber.Ctx) error {
	var m models.Author
	id := c.Params("id")
	fmt.Println(id)

	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: "+err.Error())
	}

	err = a.stg.Author.AuthorUpdate(&m, id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "update da xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusOK, "author updated successfully")
}

func (a *Api) GetAuthorList(c *fiber.Ctx) error {
	var l models.AuthorList
	err := a.stg.AuthorList.GetAuthorList(&l)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "Get qilishda xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusOK, l)
}

func (a *Api) DeleteAuthor(c *fiber.Ctx) error {
	var k models.Author
	err := c.BodyParser(&k)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "Body parcer da xatolik"+err.Error())
	}
	err = a.stg.Author.AuthorDelete(&k)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "Delete qilishda xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "Author deleted successfully")
}
