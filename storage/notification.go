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


func NewNotification(db *sql.DB) *Notification {
	return &Notification{db: db}
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

func (r *Notification) Update(id *string, n *models.Notification) error {
	res, err := r.db.Exec(`
	UPDATE 
		"notification"
	SET
		body=$2,
		type=$3
	WHERE
		id=$1
	`, id, n.Type,n.Body)
	
	if err != nil {
		return fmt.Errorf("Notification Update funcsiyada xato " + err.Error())
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("Notification Update funksiyasida xato (RowsAffected): %v", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("bunday id topilmadi:"+err.Error())
	}
	return nil
}

func (r *Notification) Delete(id *string) error {
	res, err := r.db.Exec(`
	DELETE FROM 
		"notification"
	WHERE
		id=$1
	`, id)
	if err != nil {
		return fmt.Errorf("notification Delete funcsiyada xato " + err.Error())
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("Notification Delete funksiyasida xato (RowsAffected): %v", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("bunday id  topilmadi:"+err.Error())
	}

	return nil
}

func (r *Notification) Getlist()( *models.List, error) {
	var m models.List
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
		return nil,  fmt.Errorf("queryda "+err.Error())
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println(" kanal yopilmadi", err)
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
			return  nil,err
		}
		m.Notifications = append(m.Notifications, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "notification"`).Scan(&m.Count)

	return &m, err
}
