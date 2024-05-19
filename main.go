package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/livghit/gophosaurus/services"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /", services.GetDinos)
	server.HandleFunc("POST /searchDinos", services.SearchDinos)

	port := ":3001"
	log.Default().Print("Server running at ", port)

	err := http.ListenAndServe(port, server)
	if err != nil {
		slog.Error(err.Error())
	}
}
