package main

import (
	"fmt"

	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
)

// Main.
func main() {

	opt := defoptions.NewDefOptions()
	serverRepo, err := repository.NewServerRepo(opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	// serverRepo := repository.NewRepo(opt.RepoFileName())
	s := new(server.Server)
	s.Start(serverRepo, opt)
}
