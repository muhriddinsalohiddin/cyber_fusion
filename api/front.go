package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) Front(c *fiber.Ctx) error {
	u, err := a.stg.User.GetByLogin(c.Query("login"))
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.Render("./files/a.html", u)
}
