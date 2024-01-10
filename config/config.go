package config

import "fmt"

const (
	host     = "localhost"
	pPort    = "5432"
	user     = "general"
	password = "0320"
	db       = "instagram"
)

var (
	ConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, pPort, user, password, db)
	Port = ":8080"
)
