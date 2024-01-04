package storage
import(
	"app/models"
	"database/sql"
)
type Notification struct{
	db *sql.DB
}

func NewNotification(db *sql.DB)*Notification{
	return &Notification{db:db}
}

func (r *Notification)Create(u models.Notification) string{
	return "bu xabar notification filedagi storagedan "
}
