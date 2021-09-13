package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/serialize"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

// Main.
func main() {

	opt := defoptions.NewdefOptions()
	serialize.InitSerialize(opt.RepoFileName())
	urlRepo := new(repository.URLRepo)
	serialize.ReadURLSFromFile(urlRepo.ToSet())
	repository.SerializeURLRepo = serialize.SaveURLSToFile
	s := new(server.Server)
	s.Start(urlRepo, opt)
}
