package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (r *User) Create(u *models.User) error {
	_, err := r.db.Exec(`insert into "user" (id, name, gender, birthday, email, login, password, bio) values ($1, $2,$3,$4,$5,$6,$7,$8)`,
		uuid.NewString(),
		u.Name,
		u.Gender,
		u.Birthday,
		u.Email,
		u.Login,
		u.Password,
		u.Bio,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(u1 models.User, id string) error {
	_, err := u.db.Exec(`UPDATE "user" SET name=$1, 
gender=$2, 
birthday=$3, 
email=$4, 
login=$5, 
password=$6, 
bio=$7,
updated_at=now() where id=$8`, u1.Name, u1.Gender, u1.Birthday, u1.Email, u1.Login, u1.Password, u1.Bio, id)
	if err != nil {
		return err
	}
	return nil

}

func (d *User) Delete(id string) error {
	_, err := d.db.Exec(`DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("User delete funksiyada xatolik kettiyou brat" + err.Error())
	}
	return nil
}

func (n *User) GetById(id string) (*models.User, error) {
	var updatedAt sql.NullString
	var u models.User
	err := n.db.QueryRow(`SELECT id, 
	name, 
	gender, 
	birthday, 
	email, 
	login, 
	password, 
	bio, 
	to_char(created_at, 'YYYY-MM-DD') as created_at,
	to_char(updated_at, 'YYYY-MM-DD') as updated_at 
	FROM 
	"user" 
	where id=$1`, id).Scan(
		&u.Id,
		&u.Name,
		&u.Gender,
		&u.Birthday,
		&u.Email,
		&u.Login,
		&u.Password,
		&u.Bio,
		&u.CreatedAt,
		&u.UpdetadAt)

	if updatedAt.Valid {
		u.UpdetadAt = updatedAt.String
	}

	return &u, err

}
func (n *User) Get() (*models.Users, error) {
	var (
		updatedAt sql.NullString
		m         models.Users
	)
	rows, err := n.db.Query(`SELECT id, 
	name, 
	gender, 
	birthday, 
	email, 
	login, 
	password, 
	bio, 
	to_char(created_at, 'YYYY-MM-DD') as created_at,
	to_char(updated_at, 'YYYY-MM-DD') as updated_at 
	FROM 
	"user" `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User

		err = rows.Scan(
			&u.Id,
			&u.Name,
			&u.Gender,
			&u.Birthday,
			&u.Email,
			&u.Login,
			&u.Password,
			&u.Bio,
			&u.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Get dagi rowsni scan qilishda xatolik bor ekan xatolik: " + err.Error())
		}

		if updatedAt.Valid {
			u.UpdetadAt = updatedAt.String
		} else {
			u.UpdetadAt = "bu profilda yangilanish kiritilmagan"
		}

		m.Users = append(m.Users, &u)
	}

	return &m, nil
}
