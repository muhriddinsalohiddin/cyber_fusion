package config

import "fmt"

const (
	host     = "localhost"
	pPort    = "5432"
	user     = "farrux"
	password = "1"
	db       = "farrux"
)

var (
	ConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, pPort, user, password, db)
	Port = ":8080"
)
