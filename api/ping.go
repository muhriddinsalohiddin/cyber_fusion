package api

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON("pong")
}

func PostPing(c *fiber.Ctx) error {
	return handlerResponse(c, 200, "pong post")
}

func handlerResponse(c *fiber.Ctx, code int, response any) error {
	c.SendStatus(code)
	return c.JSON(response)
}
