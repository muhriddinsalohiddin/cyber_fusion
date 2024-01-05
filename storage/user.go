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

func (u *User) Update() {

}

func (d *User) Delete() {

}

func (n *User) GetList(id string) (*models.User, error) {
	var u models.User
	row := n.db.QueryRow(`SELECT id, 
	name, 
	gender, 
	birthday, 
	email, 
	login, 
	password, 
	bio, 
	to_char(created_at, 'YYYY-MM-DD') as created_at 
	FROM 
	"user" 
	where id=$1`, id)
	row.Scan(
		&u.Id,
		&u.Name,
		&u.Gender,
		&u.Birthday,
		&u.Email,
		&u.Login,
		&u.Password,
		&u.Bio,
		&u.CreatedAt)
	// Send users as JSON response using Fiber
	fmt.Println(u)
	return &u, nil
}

// func (n *Notification) Update(u *models.User) {

// }

// func (n *User) Delete(u *models.User) {
// 	_, err := n.db.Exec(`DELETE FROM "user"
// 	WHERE id= $1;
// 	 `, u.Id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("SUCCESS")
// }
