package api

import (
	"app/models"
	"net/http"
	"strconv"

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
	id := c.Params("id")

	user, err := a.stg.User.GetById(id)
	if err != nil {
		return err
	}

	return handlerResponse(c, http.StatusCreated, user)
}

func (a *Api) GetUsers(c *fiber.Ctx) error {
	var (
		offset = 0
		limit  = 0
		err    error
	)
	limit, err = a.getLimitAndOffset(c, "limit")
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "limitni intga o'girishda xatolik 40: "+err.Error())
	}

	offset, err = a.getLimitAndOffset(c, "offset")
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "offsetni intga o'girishda xatolik 40: "+err.Error())
	}

	user, err := a.stg.User.Get(&models.UserReq{
		Limit:    limit,
		Offset:   offset,
		FromDate: c.Query("from_date", "2023-01-01"),
		ToDate:   c.Query("to_date", "2024-06-06"),
	})

	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusOK, user)
}

func (a *Api) UpdateUser(c *fiber.Ctx) error {
	var (
		u  models.User
		id = c.Params("id")
	)

	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 28: "+err.Error())
	}

	err = a.stg.User.Update(u, id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "storage"+err.Error())
	}

	return handlerResponse(c, http.StatusNoContent, "SUCCESS")
}

func (a *Api) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := a.stg.User.Delete(id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusNoContent, "TEG TUGI BILAN O'CHIP KETTI :)")
}

func (a *Api) getLimitAndOffset(c *fiber.Ctx, key string) (int, error) {
	if c.Query(key) != "" {
		return strconv.Atoi(c.Query(key))
	}
	return 0, nil
}
