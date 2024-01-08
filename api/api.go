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

type Author struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type AuthorList struct {
	Authors []*Author `json:"authors"`
	Count   int       `json:"count"`
}

func NewApi(stg *storage.Storage) *Api {

	f := fiber.New()
	a := &Api{
		f:   f,
		stg: stg,
	}

	f.Static("/images", "./images")

	f.Get("/ping", Ping)
	f.Post("/ping", PostPing)

	// user route
	{
		u := f.Group("user")
		u.Post("/", a.CreateUser)

		// u.Get("/", a.GetUser)
		// u.Get("/:id", a.GetByIdUser)
		// u.Put("/", a.UpdateUser)
		// u.Delete("/", a.DeleteUser)
	}
	{
		b := f.Group("author")
		b.Post("/", a.CreateAuthor)
		b.Put("/:id", a.UpdateAuthor)
		b.Get("/", a.GetAuthorList)
		b.Delete("/", a.DeleteAuthor)
	}

	{
		u := f.Group("post")
		u.Post("/", a.CreatePost)
		u.Delete("/",a.DeletePost)
		u.Put("/",a.UpdatePost)
		u.Get("/",a.GetPost)

	}

	return a
}

func (a *Api) Run() {
	a.f.Listen(config.Port)
}