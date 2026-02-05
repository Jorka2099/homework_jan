package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Jorka2099/homework_jan/internal/handlers"
)

// Server содержит параметр логгера и сервера
type Server struct {
	Log    *log.Logger
	Server *http.Server
}

// MakeServer создает маршрутизатор и возвращает экземпляр структуры с параметрами
func MakeServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GetHTML)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Log:    logger,
		Server: srv,
	}
}
