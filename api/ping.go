package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.JSON("salom")
}

func PostPing(c *fiber.Ctx) error {
	return handlerResponse(c, 200, "pong post")
}

func handlerResponse(c *fiber.Ctx, code int, response any) error {
	return c.Status(code).JSON(response)
}

func (a *Api) getLimitAndOffset(c *fiber.Ctx, key string) (int, error) {
	if c.Query(key) != "" {
		return strconv.Atoi(c.Query(key))
	}
	return 0, nil
}
