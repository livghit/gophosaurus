package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/livghit/gophosaurus/services"
)

func main() {
	server := http.NewServeMux()

	server.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server.HandleFunc("GET /", services.GetDinos)
	server.HandleFunc("POST /searchDinos", services.SearchDinos)
	server.HandleFunc("GET /create/dino", services.CreateNewDinoView)

	port := ":3001"
	log.Default().Print("Server running at ", port)

	err := http.ListenAndServe(port, server)
	if err != nil {
		slog.Error(err.Error())
	}
}
