package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/nlowe/aoc2025/challenge/cmd"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	return cmd.NewRootCommand().ExecuteContext(ctx)
}
