package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Author struct {
	db *sql.DB
}

type AuthorList struct {
	db *sql.DB
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
}

func NewAuthorList(db *sql.DB) *AuthorList {
	return &AuthorList{db: db}
}

func (r *Author) CreateAuthor(u *models.Author) error {
	_, err := r.db.Exec(`
		INSERT INTO "author" (
			id,
			name
		) VALUES (
			$1,$2
		)`, uuid.NewString(), u.Name)

	if err != nil {
		return fmt.Errorf("Create qilishda xatolik" + err.Error())
	}

	return nil
}

func (r *AuthorList) GetAuthorList(m *models.AuthorList) error {
	// var resp models.AuthorList
	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM "author"`
	rows, err := r.db.Query(query)
	if err != nil {
		return err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("aka kanal yopilmadi", err)
		}
	}()

	for rows.Next() {
		var (
			updated sql.NullString
			b       models.Author
		)
		err = rows.Scan(
			&b.Id,
			&b.Name,
			&b.CreatedAt,
			// &b.UpdatedAt,
			&updated,
		)
		if err != nil {
			return err
		}
		if updated.Valid {
			b.UpdatedAt = updated.String
		}

		m.Authors = append(m.Authors, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "author"`).Scan(&m.Count)

	return err
}

func (r *Author) AuthorUpdate(b *models.Author, id string) error {
	fmt.Println(id, "xaxa")
	res, err := r.db.Exec(`
        UPDATE "author" SET
            name = $2,
            created_at = $3,
            updated_at = NOW()
        WHERE id = $1`,
		id,
		b.Name, 
		b.CreatedAt,
		b.UpdatedAt,
	)

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("bunaqa idlik user yoq aka")
	}

	fmt.Println(res)

	return err
}

func (r *Author) AuthorDelete(m *models.Author) error {
	_, err := r.db.Exec(`
	DELETE FROM 
		"author" 
	WHERE 
		id = $1
	`, m.Id)

	if err != nil {
		return fmt.Errorf("Delete dan" + err.Error())
	}

	return nil
}
