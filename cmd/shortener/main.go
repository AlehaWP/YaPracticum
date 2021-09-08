package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/serialize"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

func main() {
	urlRepo := make(repository.URLRepo)
	serialize.ReadURLSFromFile(&urlRepo)
	repository.SerializeURLRepo = serialize.SaveURLSToFile
	s := new(server.Server)
	s.Start(&urlRepo)
	defer serialize.SaveURLSToFile(&urlRepo)
}
