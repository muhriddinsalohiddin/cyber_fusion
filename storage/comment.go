package storage

import (
	"app/models"
	"database/sql"
)

type Comment struct {
	db *sql.DB
}

func NewComment(db *sql.DB) *Comment {
	return &Comment{db: db}
}

func (r *Comment) Create(c models.Comment) string {
	return "bu xabar comment filedagi storagedan keldi"
}
