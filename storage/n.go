package storage

import (
	"app/models"
	"database/sql"
	"fmt"
)

type Notification struct {
	db *sql.DB
}

func NewNotification(db *sql.DB) *Notification {
	return &Notification{db: db}
}

func (n *Notification) Create(u *models.User) {
	_, err := n.db.Exec(`insert into "user" (id, name, gender, birthday, email, login, password, bio, created_at) values ($1, $2,$3,$4,$5,$6,$7,$8,now())`, u.Id,
		u.Name,
		u.Gender,
		u.Birthday,
		u.Email,
		u.Login,
		u.Password,
		u.Bio,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS")

}

func (n *Notification) Get(u *models.User) {
	rows, err := n.db.Query(`SELECT id, name, gender, birthday, email, login, password, bio, to_char(created_at, 'YYYY-MM-DD')FROM "user";
	`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(
			&u.Id,
			&u.Name,
			&u.Gender,
			u.Birthday,
			u.Email,
			u.Login,
			u.Password,
			u.Bio,
			u.CreatedAt,
		)
		if err != nil {
			fmt.Errorf("Xatolik kattiyov n.go 53 qator " + err.Error())
		}

	}
	return 
}

func (n *Notification) Update(u *models.User) {

}

func (n *Notification) Delete(u *models.User) {
	_, err := n.db.Exec(`DELETE FROM "user"
	WHERE id= $1;
	 `, u.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS")
}
