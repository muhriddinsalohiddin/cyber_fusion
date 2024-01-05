package config

import "fmt"

const (
	host     = "localhost"
	pPort    = "5433"
	user     = "postgres"
	password = "asrline"
	db       = "postgres"
)

var (
	ConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, pPort, user, password, db)
	Port = ":8080"
)
