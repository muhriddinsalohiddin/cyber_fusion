package api

import (
	"app/config"
	"app/storage"

	//"database/sql"

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

	f.Get("/ping", Ping)
	f.Post("/ping", PostPing)

	// user route
	{
		u := f.Group("user")
		u.Post("/", a.CreateUser)
		u.Get("/", a.GetUsers)
		u.Get("/:id", a.GetByIdUser)
		u.Get("/with/:id", a.GetByIdWithAllItems)
		u.Put("/:id", a.UpdateUser)
		u.Delete("/:id", a.DeleteUser)
	}

	// book route
	{
		b := f.Group("book")
		b.Post("/", a.CreateBook)
		b.Get("/", a.GetBook)
		// b.Get("/:id", a.GetBookById)
		b.Put("/:id", a.UpdateBook)
		b.Delete("/:id", a.DeleteBook)
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
		u.Delete("/", a.DeletePost)
		u.Put("/", a.UpdatePost)
		u.Get("/", a.GetPost)
		u.Get("/:id", a.GetByIdPost)

	}

	{
		u := f.Group("notification")
		u.Post("/", a.CreateNotification)
		u.Get("/", a.GetnotificationList)
		u.Delete("/:id", a.DeleteNotification)
		u.Put("/:id", a.UpdateNotification)
	}

	{
		m := f.Group("message")
		m.Post("/", a.CreateMessage)
		m.Get("/", a.GetMessageList)
		m.Put("/:id", a.UpdateMessage)
		m.Delete("/:id", a.DeleteMessage)
		m.Get("/:id", a.GetMessageById)
	}

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
	if err := a.f.Listen(config.Port); err != nil {
		panic(err)
	}
}
