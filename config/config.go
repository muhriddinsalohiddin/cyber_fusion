package config

import (
	"fmt"
	"time"
)

const (
	host     = "localhost"
	pPort    = "5432"
	user     = "postgres"
	password = "1"
	db       = "instagram"
)

var (
	ConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, pPort, user, password, db)
	Port                = ":8080"
	TOKEN_SECRET        = "secret"
	AccessTokenDuration = time.Second * 25
)
