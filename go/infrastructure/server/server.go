package server

import (
	"os"
)

type Server interface {
	Run()
}

func NewServer() Server {
	c := newConfig()
	return NewGinServer(c)
}

type config struct {
	addr string
}

const defaultPort = "8080"

func newConfig() *config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	return &config{
		addr: ":" + port,
	}
}
