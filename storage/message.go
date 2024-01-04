package storage

import (
	"app/models"
	"database/sql"
)

type Message struct {
	db *sql.DB
}

func NewMessage(db *sql.DB) *Message {
	return &Message{db: db}
}

func (r *Message) Create(m models.Message) string {
	return "bu xabar Message filedagi storagedan keldi"
}
