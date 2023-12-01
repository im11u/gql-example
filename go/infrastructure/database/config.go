package database

import (
	"os"
)

type config struct {
	host     string
	database string
	port     string
	user     string
	password string
}

func newConfig() *config {
	return &config{
		host:     os.Getenv("DB_HOST"),
		database: os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
	}
}
