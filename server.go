package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Salaton/tracing/pkg/presentation"
)

func main() {
	presentation.StartServer(8080)

	// Block until we receive a sigint (CTRL+C) signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
}
