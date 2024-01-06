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
	return nil
}

func (r *Message) Update(m *models.Message, id *string) error {
	_, err := r.db.Exec(`
	UPDATE 
		"message"
	SET
		body=$2
	WHERE
		id=$1
	`, id, m.Body)
	if err != nil {
		return fmt.Errorf("Message Update funcsiyada xato bor akaxon" + err.Error())
	}

	return nil
}

func (r *Message) Delete(id *string) error {
	res, err := r.db.Exec(`
	DELETE FROM 
		"message"
	WHERE
		id=$1
	`, id)
	if err != nil {
		return fmt.Errorf("Message Delete funcsiyada xato bor akaxon" + err.Error())
	}
	i, err := res.RowsAffected()
	if i == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *Message) GetMessageList(req *models.ListMessageReq) (*models.ListMessage, error) {
	var (
		m     models.ListMessage
		query = `
			SELECT
				id,sender_id,receiver_id,body,created_at 
			FROM
				"message"`
		filter = " WHERE 1=1 "
		args   []any
	) 
	if req.FromDate != "" && req.ToDate!=""{
		args = append(args, req.FromDate)
		filter += " AND created_at > $" + fmt.Sprint(len(args))
		args = append(args, req.ToDate)
		filter += " AND created_at < $" + fmt.Sprint(len(args))
	}

	if req.ReceiverId != "" {
		args = append(args, req.ReceiverId)
		filter += " AND receiver_id = $" + fmt.Sprint(len(args))
	}

	if req.SenderId != "" {
		args = append(args, req.SenderId)
		filter += " AND sender_id = $" + fmt.Sprint(len(args))
	}
	query = query + filter
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Okaxon GetMessageListdagi quericha sal xato ishlayabdi " + err.Error())
	}
	defer rows.Close()
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
			return nil, fmt.Errorf("Okaxon GetMessageListdagi FORdagi Scan sal xato ishlayabdi " + err.Error())
		}
		m.Messages = append(m.Messages, &message)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "message"`+filter, args...).Scan(&m.Cout)

	return &m, err
}

func (r *Message) GetMessage(id *string) (*models.Message, error) {
	var m models.Message
	err := r.db.QueryRow(`
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
		return nil, fmt.Errorf("Message Update funcsiyada xato bor akaxon" + err.Error())
	}
	return &m, nil
}
