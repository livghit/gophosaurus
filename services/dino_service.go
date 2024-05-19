package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/livghit/gophosaurus/types"
	"github.com/livghit/gophosaurus/views"
	"github.com/livghit/gophosaurus/views/partials"
)

func GetDinos(w http.ResponseWriter, r *http.Request) {
	dinos, err := fetchDinosaurs()
	if err != nil {
		slog.Error(err.Error())
	}
	views.Index(dinos).Render(r.Context(), w)
}

func SearchDinos(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		dinos, err := fetchDinosaurs()
		if err != nil {
			slog.Error(err.Error())
		}
		formErr := r.ParseForm()
		if formErr != nil {
			slog.Error("Problem while parsing)")
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		searchQuery := strings.ToLower(r.FormValue("search"))
		var filtredDinos []types.Dino
		for _, dino := range dinos {
			if strings.Contains(strings.ToLower(dino.Name), searchQuery) {
				filtredDinos = append(filtredDinos, dino)
			}
		}
		partials.DinoCard(filtredDinos).Render(r.Context(), w)
	} else {
		w.Write([]byte("Unsuported Request Method"))
	}
}

func fetchDinosaurs() ([]types.Dino, error) {
	url := "https://dinosaur-facts-api.shultzlab.com/dinosaurs"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("Error: unexpected status code %d", response.StatusCode)
		return nil, err
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
