package api

import (
	"app/models"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CraeteComment(c *fiber.Ctx) error {
	var com models.Comment
	data := a.stg.Comment.Create(com)
	return handlerResponse(c, 201, data)
}
