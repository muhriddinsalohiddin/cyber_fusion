package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db           *sql.DB
	User         *User
	Comment      *Comment
	Message      *Message
	Notification *Notification
	Author       *Author
	Post         *Post
}

func NewStorage(connStr string) *Storage {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	return &Storage{
		db:           db,
		User:         NewUser(db),
		Comment:      NewComment(db),
		Message:      NewMessage(db),
		Notification: NewNotification(db),
		Author:       NewAuthor(db),
		Post:         NewPost(db),
	}
}

func (s *Storage) Close() {
	s.db.Close()
}
