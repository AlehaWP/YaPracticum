package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
	"github.com/AlehaWP/YaPracticum.git/internal/signal"
)

// Main.
func main() {

	mainCtx, cancel := context.WithCancel(context.Background())
	opt := defoptions.NewDefOptions()
	sr, err := repository.NewServerRepo(mainCtx, opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	defer sr.Close()
	// serverRepo := repository.NewRepo(opt.RepoFileName())
	s := new(server.Server)
	go signal.HandleQuit(cancel)
	go s.Start(mainCtx, sr, opt)

	<-mainCtx.Done()

	timer := time.NewTicker(5 * time.Second)
	<-timer.C
}
