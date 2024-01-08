package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db   *sql.DB
	User *User
	Books *Books
	// BooksList *BooksList
}

func NewStorage(connStr string) *Storage {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Storage{
		db:   db,
		User: NewUser(db),
		Books : NewBooks(db),
		// BooksList: NewBooksList(db),
	}
}

func (s *Storage) Close() {
	s.db.Close()
}
