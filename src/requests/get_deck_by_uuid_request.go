package requests

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"deck-api/src/models"
	"deck-api/src/repositories"
)

func GetDeckByUUID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var deckRepository repositories.DeckRepository
	decks := deckRepository.GetDecks()

	uuid := mux.Vars(r)["uuid"]
	idx := idxByUUID(decks, uuid)

	if len(uuid) == 0 {
		http.Error(w, "Please provide UUID", http.StatusBadRequest)
		return
	}

	if idx < 0 {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(decks[idx]); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func idxByUUID(decks []models.Deck, uuid string) int {
	for i := range decks {
		if decks[i].UUID == uuid {
			return i
		}
	}

	return -1
}
