package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
	"github.com/AlehaWP/YaPracticum.git/internal/grcpserver"
	"github.com/AlehaWP/YaPracticum.git/internal/repository"
	"github.com/AlehaWP/YaPracticum.git/internal/server"
	"github.com/AlehaWP/YaPracticum.git/internal/signal"
)

var (
	BuildVersion string = "N/A"
	BuildDate    string = "N/A"
	BuildCommit  string = "N/A"
)

// Main.
func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", BuildVersion, BuildDate, BuildCommit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opt := defoptions.NewDefOptions()

	sr, err := repository.NewServerRepo(ctx, opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	defer sr.Close()

	server := new(server.Server)
	go server.Start(ctx, sr, opt)
	go grcpserver.Start(ctx, sr, opt)

	go signal.HandleQuit(cancel)
	<-ctx.Done()
}
