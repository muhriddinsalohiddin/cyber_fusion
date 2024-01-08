package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Comment struct {
	db *sql.DB
}
type List struct {
	db *sql.DB
}

func NewComment(db *sql.DB) *Comment {
	return &Comment{db: db}
}
func NewList(db *sql.DB) *List {
	return &List{db: db}
}
func (c *Comment) Create(n *models.Comment) error {
	_, err := c.db.Exec(
		`
		INSERT INTO
				"comment"(
					id,user_id,post_id,body
				) VALUES (
					$1 ,$2, $3, $4
				)
		`, uuid.NewString(), n.UserId, n.PostId, n.Body)
	if err != nil {
		return fmt.Errorf("Comment create funksiyada xato" + err.Error())
	}
	return nil
}
func (c *Comment) Update(n *models.Comment, id *string) error {
	_, err := c.db.Exec(
		`
		UPDATE
			"comment"
		SET 
			body=$2
			WHERE
			id=$1	
			`, id, n.Body)
	if err != nil {
		return fmt.Errorf("Comment Update qilishda xato" + err.Error())
	}
	
	return nil
}
func (c *Comment) DeleteComment( id *string)  error {
	_, err := c.db.Exec(
		`
		DELETE FROM
			"comment"
		WHERE 
			id=$1
		`, id)
	if err != nil {
		return fmt.Errorf("Comment Update qilishda xato" + err.Error())
	}
	
	return nil
}
func (r *Comment) Getlist(req *models.Comment) (*models.Listcha, error) {
	var( m models.Listcha
	query = `
	SELECT
		id,
		user_id,
		post_id,
		body,
		created_at,
		updated_at
	FROM "comment"`
    filter=" Where 1=1 "
    args  []any)
	if req.UserId !=""{
		args=append(args,req.UserId )
		filter +=" AND user_id=$"+fmt.Sprint(len(args))
	}
	query=query+filter
	rows, err := r.db.Query(query,args...)
	if err != nil {
		return nil, fmt.Errorf("queryda " + err.Error())
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println(" kanal yopilmadi", err)
		}
	}()

	for rows.Next() {
		var (
			c models.Comment
			updated sql.NullString
		)
		err = rows.Scan(
			&c.Id,
			&c.UserId,
			&c.PostId,
			&c.Body,
			&c.CreatedAt,
			&updated,
		)
		if err != nil {
			return nil, err
		}
		if updated.Valid {
			c.UpdatedAt = updated.String
		}
		m.Comments = append(m.Comments, &c)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "comment"`+filter, args...).Scan(&m.Count)

	return &m, err
}