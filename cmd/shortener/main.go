package main

import (
	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

// Main.
func main() {

	opt := defoptions.NewDefOptions()
	serverRepo := repository.NewRepo(opt.RepoFileName())
	s := new(server.Server)
	s.Start(serverRepo, opt)
}
