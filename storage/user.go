package storage

import (
	"app/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (r *User) Create(u models.User) string {
	return "bu xabar user filedagi storagedan keldi"
}
