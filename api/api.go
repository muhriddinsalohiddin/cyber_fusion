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
	//db *sql.DB
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
		// u.Get("/", a.GetUser)
		// u.Get("/:id", a.GetByIdUser)
		// u.Put("/", a.UpdateUser)
		// u.Delete("/", a.DeleteUser)
	}
	{
		u:=f.Group("notification")
		u.Post("/",a.CreateNotification)
		u.Get("/",a.GetnotificationList)
		u.Delete("/:id",a.DeleteNotification)
		u.Put("/:id",a.UpdateNotification)
	}

	return a
}

func (a *Api) Run() {
	a.f.Listen(config.Port)
}
