package server

import (
	"os"
)

type Server interface {
	Run()
}

func NewServer() Server {
	c := newConfig()
	return &GinServer{config: c}
}

type config struct {
	addr string
}

func newConfig() *config {
	port := os.Getenv("APP_PORT")

	return &config{
		addr: ":" + port,
	}
}
