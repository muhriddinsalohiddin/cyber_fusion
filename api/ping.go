package api

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON("salom")
}

func PostPing(c *fiber.Ctx) error {
	return handlerResponse(c, 200, "pong post")
}

func handlerResponse(c *fiber.Ctx, code int, response any) error {
	return c.JSON(response)
}
