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

func (r *Author) Create(u models.Author) string {
	return "bu xabar author faylidagi storage dan keldi"
}

func CreateAuthor(db *sql.DB, u *models.Author) error {
	_, err := db.Exec(`
	INSERT INTO "user" (
		id,
		name,
		created_at,
		updated_at
	) VALUES (
		$1,$2,$3,$4
	)`, uuid.NewString(), u.Name, u.CreatedAt, u.UpdatedAt)

	return err
}

// func GetUserList(db *sql.DB) (*models.Author, error) {
// 	var resp models.AuthorList
// 	query := `
// 	SELECT
// 		id,
// 		name,
// 		to_char(created_at,'DD.MM.YYYY HH24:MI:SS'),
// 		to_char(updated_at,'DD.MM.YYYY HH24:MI:SS')
// 	FROM "author"`
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		err = rows.Close()
// 		if err != nil {
// 			fmt.Println("aka kanal yopilmadi", err)
// 		}
// 	}()

// 	for rows.Next() {
// 		var (
// 			updated sql.NullString
// 			u       models.User
// 		)
// 		err = rows.Scan(
// 			&u.Id,
// 			&u.Name,
// 			&u.CreatedAt,
// 			&u.UpdatedAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if updated.Valid {
// 			u.UpdatedAt = updated.String
// 		}
// 		resp.Authors = append(resp.Authors, &(*models.Author)(u))
// 	}

// 	err = db.QueryRow(`SELECT COUNT(1) FROM "author"`).
// 		Scan(&resp.Count)

// 	return &resp, err
// }

func AuthorUpdate(db *sql.DB, u *models.Author) error {
	res, err := db.Exec(`
	UPDATE "author" SET
		name = $2
		created_at = $3
		updated_at = NOW()
		WHERE id = $1`,
		u.Id,
		u.Name,
		u.CreatedAt,
		u.UpdatedAt,
	)

	if rows, _ := res.RowsAffected(); rows == 0 {
		return fmt.Errorf("bunaqa idlik user yoq aka")
	}

	return err
}

func AuthorDelete(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM \"user\" WHERE id = $1", id)

	if row, _ := res.RowsAffected(); row == 0 {
		return fmt.Errorf("bunaqa idlik user yo'q")
	}

	return err
}
