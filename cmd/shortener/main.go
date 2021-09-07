package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/projectenv"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/serialize"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

func main() {
	projectenv.Init()
	err := serialize.Init(projectenv.Envs.OptionsFileName)
	if err != nil {
		return
	}

	urlRepo := make(repository.URLRepo)
	serialize.ReadURLSFromFile(&urlRepo)
	s := new(server.Server)
	s.Start(&urlRepo)
	defer serialize.SaveURLSToFile(urlRepo)

}
