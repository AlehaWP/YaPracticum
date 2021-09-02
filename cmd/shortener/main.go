package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

func main() {
	s := new(server.Server)
	urlRepo := make(repository.URLRepo)
	s.Start("localhost:8080", &urlRepo)
}
