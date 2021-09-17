package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/serialize"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

// Main.
func main() {

	opt := defoptions.NewDefOptions()

	urlRepo := repository.NewUrlRepo()

	serialize.NewSerialize(opt.RepoFileName())
	serialize.ReadURLSFromFile(urlRepo)
	repository.SerializeURLRepo = serialize.SaveURLSToFile

	s := new(server.Server)
	s.Start(urlRepo, opt)
}
