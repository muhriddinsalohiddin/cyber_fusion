package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Post struct {
	db *sql.DB
}
type List struct{
	db *sql.DB
}

func NewPost(db *sql.DB) *Post {
	return &Post{db: db}
}
func NewList(db *sql.DB) *List{
	return &List{db:db}
}

func (r *Post) Create(m *models.Post) error {
	_, err := r.db.Exec(`
		INSERT INTO 
			"post"(
				id,user_id,title,body
			) 
		VALUES(
			$1,$2,$3,$4
		)
	`,uuid.NewString(),m.UserId,m.Title,m.Body)
	if err != nil {
		return fmt.Errorf("Post create funksiya xato bor akaxon"+err.Error())
	}
	return nil
}


func(r *Post) Delete(m *models.Post) error{
	_, err := r.db.Exec(`
	DELETE FROM 
		"post"
	WHERE
		id=$1
	`,m.Id)
	if err != nil {
		return fmt.Errorf("Post Delete funcsiyada xato bor akaxon"+err.Error())
	}
	return nil
}
func (r *Post) Update(n *models.Post) error {
	_, err := r.db.Exec(`
	UPDATE 
		"post"
	SET
		body=$2,
		title=$3
	WHERE
		id=$1
	`, n.Id, n.Body,n.Title)
	if err != nil {
		return fmt.Errorf("Post Update funksiya xato " + err.Error())
	}
	return nil
}
func (r *List) GetPostlist(m *models.List)  error {
	query := `
	SELECT
		id,
		user_id,
		title,
		body,
		created_at
	FROM "post"`
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
			b       models.Post
		)
		err = rows.Scan(
			&b.Id,
			&b.UserId,
			&b.Title,
			&b.Body,
			&b.CreatedAt,
		)
		if err != nil {
			return  err
		}
		m.Post = append(m.Post, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "post"`).Scan(&m.Cout)

	return  err
}