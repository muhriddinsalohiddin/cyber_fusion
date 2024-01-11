package api

import (
	"app/config"
	"app/helper"
	"app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateUser(c *fiber.Ctx) error {
	var u models.User

	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 15: "+err.Error())
	}
	u.Password, _ = helper.HashPassword(u.Password)

	err = a.stg.User.Create(&u)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return handlerResponse(c, http.StatusCreated, "SUCCESS")
}

func (a *Api) GetByIdWithAllItems(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := a.stg.User.GetByIdWithAllItems(id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "get user"+err.Error())
	}

	return handlerResponse(c, http.StatusCreated, user)
}
func (a *Api) GetByIdUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := a.stg.User.GetById(id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "get user"+err.Error())
	}

	comments, err := a.stg.Comment.Getlist(&models.Comment{
		UserId: user.Id,
	})
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "get comment"+err.Error())
	}

	user.Comments = comments.Comments

	notifications, err := a.stg.Notification.Getlist(&models.Notification{
		UserId: user.Id,
	})
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "get notification"+err.Error())
	}
	user.Notifications = notifications.Notifications

	posts, err := a.stg.Post.GetlistWithComments(user.Id)

	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "get post"+err.Error())
	}
	user.Posts = posts.Post

	return handlerResponse(c, http.StatusCreated, user)
}

func (a *Api) GetUsers(c *fiber.Ctx) error {
	var (
		offset = 0
		limit  = 0
		err    error
	)

	ok, err := a.checkAuth(c)
	if !ok {
		return handlerResponse(c, http.StatusUnauthorized, err.Error())
	}

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
		Login:    c.Query("login"),
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

func (a *Api) Login(c *fiber.Ctx) error {
	var (
		u models.Login
	)

	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik 28: "+err.Error())
	}

	user, err := a.stg.User.GetByLogin(u.Login)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "storage"+err.Error())
	}

	if !helper.CheckPasswordHash(u.Password, user.Password) {
		return handlerResponse(c, http.StatusBadRequest, "login yoki parol xato")
	}

	token, err := helper.NewAccessToken(models.UserClaims{
		Id:       user.Id,
		Name:     user.Name,
		Duration: time.Now().Add(config.AccessTokenDuration).Unix(),
		Login:    user.Login,
	})
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "token"+err.Error())
	}

	user.Token = token
	return handlerResponse(c, http.StatusOK, user)
}
func (a *Api) checkAuth(c *fiber.Ctx) (bool, error) {
	token := c.Get("token")

	s, err := helper.ParseAccessToken(token)
	if err != nil {
		return false, err
	}

	if s.Duration < time.Now().Unix() {
		return false, fmt.Errorf("token muddati tugagan")
	}
	return true, nil
}
