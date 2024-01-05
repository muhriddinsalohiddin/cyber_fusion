package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Notification struct {
	db *sql.DB
}
type List struct{
	db *sql.DB
}

func NewNotification(db *sql.DB) *Notification {
	return &Notification{db: db}
}
func NewList(db *sql.DB) *List{
	return &List{db:db}
}

func (r *Notification) Create(n *models.Notification) error {
	_, err := r.db.Exec(`
		INSERT INTO 
			"notification"(
				id,user_id,type,body
			) 
		VALUES(
			$1,$2,$3,$4
		)
	`, uuid.NewString(), n.UserId, n.Type, n.Body)
	if err != nil {
		return fmt.Errorf("notification create funcsiyada xato " + err.Error())
	}
	return nil
}

func (r *Notification) Update(n *models.Notification) error {
	_, err := r.db.Exec(`
	UPDATE 
		"notification"
	SET
		body=$2,
		type=$3
	WHERE
		id=$1
	`, n.Id, n.Type,n.Body)
	if err != nil {
		return fmt.Errorf("Notification Update funcsiyada xato " + err.Error())
	}
	return nil
}

func (r *Notification) Delete(n *models.Notification) error {
	_, err := r.db.Exec(`
	DELETE FROM 
		"notification"
	WHERE
		id=$1
	`, n.Id)
	if err != nil {
		return fmt.Errorf("notification Delete funcsiyada xato " + err.Error())
	}
	return nil
}

func (r *List) Getlist(m *models.List)  error {
	query := `
	SELECT
		id,
		user_id,
		type,
		body,
		created_at
	FROM "notification"`
	rows, err := r.db.Query(query)
	if err != nil {
		return  fmt.Errorf("queryda ")
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("aka kanal yopilmadi", err)
		}
	}()

	for rows.Next() {
		var (
			b       models.Notification
		)
		err = rows.Scan(
			&b.Id,
			&b.UserId,
			&b.Type,
			&b.Body,
			&b.CreatedAt,
		)
		if err != nil {
			return  err
		}
		m.Notifications = append(m.Notifications, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "notification"`).Scan(&m.Count)

	return  err
}
