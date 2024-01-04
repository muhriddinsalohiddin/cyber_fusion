package api
import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateNotification(c *fiber.Ctx) error {
	var u models.Notification
	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parcerda xatolik: "+err.Error())
	}
	fmt.Println(u)
	return handlerResponse(c, http.StatusCreated, u)
	
}