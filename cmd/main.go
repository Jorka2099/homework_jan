package main

import (
	"log"
	"os"

	"github.com/Jorka2099/homework_jan/internal/server"
)

// main создает сервер и слушает порт
func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.MakeServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal("Failed to start server ", err)
	}
}
