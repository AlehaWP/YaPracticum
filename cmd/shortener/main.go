package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	"github.com/AlehaWP/YaPracticum.git/internal/defoptions"
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

	server := new(server.Server)
	go signal.HandleQuit(cancel)
	go server.Start(ctx, opt)
	//go grcpserver.Start(ctx, opt)
	<-ctx.Done()
}
