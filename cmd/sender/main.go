package main

import (
	"context"
	"github.com/ra9dev/go-e/internal/app/sender"
	"github.com/ra9dev/go-e/internal/pkg/config"
	"github.com/ra9dev/go-e/internal/pkg/os"
)

func main() {
	cfg := new(sender.TimeSenderConfig)
	config.ParseConfig(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go os.CatchTermination(cancel)

	server := sender.NewTimeSender(cfg)
	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
