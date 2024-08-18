package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mark-marushak/microservice/apps/online-shop/cmd/serve"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "Auth Service",
		Usage:   "Run application to Auth Service working",
		Authors: []any{"markk.9964@gmail.com"},
		Commands: []*cli.Command{
			serve.New(),
		},
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatalf("run cli app: %v", err)
	}
}
