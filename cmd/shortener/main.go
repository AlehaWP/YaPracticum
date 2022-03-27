package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
	"github.com/AlehaWP/YaPracticum.git/internal/signal"
)

// Main.
func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opt := defoptions.NewDefOptions()
	sr, err := repository.NewServerRepo(ctx, opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	defer sr.Close()
	// serverRepo := repository.NewRepo(opt.RepoFileName())
	s := new(server.Server)
	go signal.HandleQuit(cancel)
	go s.Start(ctx, sr, opt)

	<-ctx.Done()
}
