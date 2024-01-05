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
type List struct {
	db *sql.DB
}

func NewMessage(db *sql.DB) *Message {
	return &Message{db: db}
}
func NewList(db *sql.DB) *List {
	return &List{db: db}
}

func (r *Message) Create(m *models.Message) error {
	_, err := r.db.Exec(`
		INSERT INTO 
			"message"(
				id,sender_id,receiver_id,body
			) 
		VALUES(
			$1,$2,$3,$4
		)
	`, uuid.NewString(), m.SenderId, m.ReceiverId, m.Body)
	if err != nil {
		return fmt.Errorf("Message create funcsiyada xato bor akaxon" + err.Error())
	}
	return err
}

func (r *Message) Update(m *models.Message) error {
	_, err := r.db.Exec(`
	UPDATE 
		"message"
	SET
		body=$2
	WHERE
		id=$1
	`, m.Id, m.Body)
	if err != nil {
		return fmt.Errorf("Message Update funcsiyada xato bor akaxon" + err.Error())
	}
	return err
}

func (r *Message) Delete(id *string) error {
	_, err := r.db.Exec(`
	DELETE FROM 
		"message"
	WHERE
		id=$1
	`, id)
	if err != nil {
		return fmt.Errorf("Message Delete funcsiyada xato bor akaxon" + err.Error())
	}
	return err
}

func (r *List) GetMessageList(m *models.List) error {
	query := `
		SELECT
			id,sender_id,receiver_id,body,created_at 
		FROM
			"message"
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return fmt.Errorf("Okaxon GetMessageListdagi quericha sal xato ishlayabdi " + err.Error())
	}
	for rows.Next() {
		var message models.Message
		err = rows.Scan(
			&message.Id,
			&message.SenderId,
			&message.ReceiverId,
			&message.Body,
			&message.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("Okaxon GetMessageListdagi FORdagi Scan sal xato ishlayabdi " + err.Error())
		}
		m.Messages = append(m.Messages, &message)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "message"`).Scan(&m.Cout)

	return err
}

func (r *Message) GetMessage(m *models.Message, id *string) error {
	err:= r.db.QueryRow(`
	SELECT 
		id,sender_id,receiver_id,body,created_at
	FROM
		"message"
	WHERE
		id=$1
	`, id).Scan(
		&m.Id,
		&m.SenderId,
		&m.ReceiverId,
		&m.Body,
		&m.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("Message Update funcsiyada xato bor akaxon" + err.Error())
	}
	return nil
}
