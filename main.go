package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/livghit/gophosaurus/types"
	"github.com/livghit/gophosaurus/views"
)

func main() {
	server := http.NewServeMux()
	dinos, err := fetchDinosaurs()
	if err != nil {
		return
	}

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		views.Hello(dinos).Render(r.Context(), w)
	})

	port := ":3000"
	log.Default().Print("Server running at ", port)
	http.ListenAndServe(port, server)
}

func fetchDinosaurs() ([]types.Dino, error) {
	url := "https://dinosaur-facts-api.shultzlab.com/dinosaurs"

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching dinosaurs data: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: unexpected status code %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	var dinos []types.Dino
	err = json.Unmarshal(data, &dinos)
	if err != nil {
		panic(err)
	}
	return dinos, nil
}
