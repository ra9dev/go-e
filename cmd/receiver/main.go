package main

import (
	"context"
	"github.com/ra9dev/go-e/internal/app/receiver"
	"github.com/ra9dev/go-e/internal/pkg/config"
	"github.com/ra9dev/go-e/internal/pkg/os"
)

func main() {
	cfg := new(receiver.TimeConsumerConfig)
	config.ParseConfig(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go os.CatchTermination(cancel)

	consumer := receiver.NewTimeConsumer(cfg)
	if err := consumer.Run(ctx); err != nil {
		panic(err)
	}
}
