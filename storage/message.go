package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Message struct {
	db *sql.DB
}

func NewMessage(db *sql.DB) *Message {
	return &Message{db: db}
}

func (r *Message) Create(m models.Message) error {
	_, err := r.db.Exec(`
		INSERT INTO 
			"message"(
				id,sender_id,receiver_id,body
			) 
		VALUES(
			$1,$2,$3,$4
		)
	`,uuid.NewString(),m.SenderId,m.ReceiverId,m.Body)
	if err != nil {
		return fmt.Errorf("Message create funcsiyada xato bor akaxon"+err.Error())
	}
	return nil
}

func (r *Message) Update(m models.Message) error {
	_, err := r.db.Exec(`
	UPDATE 
		"message"
	SET
		body=$2
	WHERE
		id=$1
	`,m.Id,m.Body)
	if err != nil {
		return fmt.Errorf("Message Update funcsiyada xato bor akaxon"+err.Error())
	}
	return nil
}
