package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"

	"github.com/livghit/gophosaurus/types"
	"github.com/livghit/gophosaurus/views"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		dinos, err := fetchDinosaurs()
		if err != nil {
			slog.Error(err.Error())
		}
		views.DinoCards(dinos).Render(r.Context(), w)
	})

	port := ":3000"
	log.Default().Print("Server running at ", port)
	http.ListenAndServe(port, server)
}

func fetchDinosaurs() ([]types.Dino, error) {
	url := "https://dinosaur-facts-api.shultzlab.com/dinosaurs"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Errorf("Error: unexpected status code %d", response.StatusCode)
		return nil, nil
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var dinos []types.Dino
	err = json.Unmarshal(data, &dinos)
	if err != nil {
		panic(err)
	}
	return dinos, nil
}
