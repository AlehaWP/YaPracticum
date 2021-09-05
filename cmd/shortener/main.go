package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/projectenv"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

func main() {
	projectenv.Init()
	s := new(server.Server)
	urlRepo := make(repository.URLRepo)
	s.Start(&urlRepo)
}
