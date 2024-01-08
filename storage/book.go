package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Books struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
	return &Books{db: db}
}

func (r *Books) CreateBook(b *models.Books) error {
	_, err := r.db.Exec(`
	INSERT INTO "book"(
		id,
		 title,
		  author,
			 description
			) VALUES (
				$1,$2,$3,$4
	)`, uuid.NewString(), b.Title, b.Author, b.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Books) GetList(req models.LsitBookReq) (*models.BooksList, error) {
	// var resp models.BooksList
	var (
		bs    models.BooksList
		query = ` select
	id,
	title,
 	author,
	description,
	created_at,
	updated_at
	FROM "book"`
		filter = "Where 1=1"
		args   []any
	)

	if req.Author != "" {
		args = append(args, req.Author)
		filter += " And author = $" + fmt.Sprint(len(args))
	}

	if req.Title != "" {
		args = append(args, req.Title)
		filter += "And title = $" + fmt.Sprint(len(args))
	}

	// if req.Id != "" {
	// 	args = append(args, req.Id)
	// 	filter
	// }

	query = query + filter
	fmt.Println("query", query)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Dear Friend in GetList have error" + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var (
			updated sql.NullString
			b      models.Books
		)

		err = rows.Scan(
			&b.Id,
			&b.Title,
			&b.Author,
			&b.Description,
			&b.CreatedAt,
			&updated,
		)
		if err != nil {
			return nil,  fmt.Errorf("Dear Friend in getList for scan don't word" + err.Error())
		}
		bs.Books =append(bs.Books, &b)
	}
	err = r.db.QueryRow(`Select Count(1) from "book"` + filter, args...).Scan(&bs.Count)
	return &bs, err
}

func (r *Books) UpdateBook(b *models.Books, id *string) error {
	_, err := r.db.Exec(`
UPDATE "book" SET 
title = $2,
 author = $3,
  description = $4
	where id = $1`,
		id,
		b.Title,
		b.Author,
		b.Description,
	)
	if err != nil {
		return  fmt.Errorf("Updated at da parcerda xatolik: "+err.Error())
	}
	return nil
}

func (r *Books) DeleteBook(id *string) error {
	_, err := r.db.Exec(`
	Delete From 
		"book"
		where id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("The delete function don't work" + err.Error())
	}
	return err
}

func (r *Books) GetBookById(b *models.Books, id *string) error {
	err := r.db.QueryRow(`
	SELECT 
	id, title, author, description, created_at 
	FROM "book" 
	where id = $1
	`, id).Scan(
		&b.Id,
		&b.Title,
		&b.Author,
		&b.Description,
		&b.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("In GetBookById function don't work")
	}
	return nil
}
