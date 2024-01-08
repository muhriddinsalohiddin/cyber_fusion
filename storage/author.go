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

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
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

func (r *Author) GetList() (*models.AuthorList, error) {
	// var m *models.Author
	var res models.AuthorList
	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM "author"`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
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
			&updated,
		)
		if err != nil {
			return nil, err
		}
		if updated.Valid {
			b.UpdatedAt = updated.String
		}

		res.Authors = append(res.Authors, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "author"`).
		Scan(&res.Count)

	return &res, err
}

func (r *Author) Update(b *models.Author, id *string) error {
	res, err := r.db.Exec(`
        UPDATE 
			"author" 
		SET
            name = $2
        WHERE 
			id = $1`,
		id,
		b.Name,
	)

	if rowsAffected, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("Update funksiyasida xato (RowsAffected): %v", err)
	} else if rowsAffected == 0 {
		return fmt.Errorf("bunday id topilmadi:" + err.Error())
	}

	return err
}

func (r *Author) Delete(m *models.Author) error {
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
