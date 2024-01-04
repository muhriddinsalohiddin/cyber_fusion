package storage

import (
	"app/models"
	"database/sql"
)

type Author struct {
	db *sql.DB
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db : db}
}

func (r *Author)Create(u models.Author) string {
	return "bu xabar author faylidagi storage dan keldi"
}