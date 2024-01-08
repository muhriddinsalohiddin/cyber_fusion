package api

import (
	"app/config"
	"app/storage"

	"github.com/gofiber/fiber/v2"
)

type Api struct {
	f   *fiber.App
	stg *storage.Storage
}

func NewApi(stg *storage.Storage) *Api {

	f := fiber.New()
	a := &Api{
		f:   f,
		stg: stg,
	}

	
	// user route
	// {
	// 	u := f.Group("user")
	// 	u.Post("/", a.CreateUser)
	// 	// u.Get("/", a.GetUser)
	// 	//  u.Get("/:id", a.GetByIdUser)
	// 	//  u.Put("/", a.UpdateUser)
	// 	//  u.Delete("/", a.DeleteUser)
	// }

	{
		c := f.Group("comments")
		c.Post("/", a.CreateComment)
		c.Get("/", a.GetCommentlist)
		c.Delete("/:id", a.DeleteComment)
		c.Put("/:id", a.UpdateComment)
	}
	return a

}

func (a *Api) Run() {
	a.f.Listen(config.Port)
}
