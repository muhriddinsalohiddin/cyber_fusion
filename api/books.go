package api

import (
	"app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateBook(c *fiber.Ctx) error {
	var b models.Books
	err := c.BodyParser(&b)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	err = a.stg.Books.CreateBook(&b)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusCreated, b)
}

func (a *Api) GetBook(c *fiber.Ctx) error {
	b, err := a.stg.Books.GetList(models.LsitBookReq{
		Author: c.Query("author"),
		Title: c.Query("title"),
		
	})
	if  err != nil {
		return handlerResponse(c, http.StatusBadRequest, "GetBook parcerda xatolik: "+err.Error())
	}
	return handlerResponse(c, http.StatusOK, b)
}

// func (a *Api) GetBookById(c *fiber.Ctx) error {
// 	var b *models.Books
// 	id := c.Params("id")
// 	a.stg.Books.GetBookById(b, &id)
// 	return handlerResponse(c, http.StatusOK, "I send you your id"+id)
// }

func (a *Api) UpdateBook(c *fiber.Ctx) error {
	var b models.Books
	err := c.BodyParser(&b)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "Update parcerda xatolik: "+err.Error())
	}
	id := c.Params("id")
	err = a.stg.Books.UpdateBook(&b, &id)
	if err != nil {
		 return handlerResponse(c, http.StatusBadRequest, "Update parcerda xatolik: "+err.Error())
	}
	return handlerResponse(c, http.StatusOK, "Updated")
}

func (a *Api) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	a.stg.Books.DeleteBook(&id)
	return handlerResponse(c, http.StatusOK, nil)
}
