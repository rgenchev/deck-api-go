package requests

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"deck-api/src/models"
	"deck-api/src/repositories"
	"deck-api/src/responses"
)

func DrawCardsFromDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var deckRepository repositories.DeckRepository
	decks := deckRepository.GetDecks()

	uuid := mux.Vars(r)["uuid"]
	numOfDrawnCardsParam := mux.Vars(r)["count"]
	numOfDrawnCards, _ := strconv.Atoi(numOfDrawnCardsParam)
	idx := idxByUUID(decks, uuid)

	if idx < 0 {
		http.Error(w, "Deck not found", http.StatusNotFound)
		return
	}

	deck := decks[idx]

	if len(deck.Cards) < numOfDrawnCards {
		http.Error(w, "Not enough cards in the deck. Please check the number of cards remaining and modify the count parameter.", http.StatusBadRequest)
		return
	}

	drawnCards := getDrawnCards(deck, numOfDrawnCards)

	deck.Cards = deck.Cards[:len(deck.Cards)-numOfDrawnCards]
	deck.Remaining = deck.Remaining - len(drawnCards)

	decks[idx] = deck

	drawCardsFromDeckResponse := buildDrawCardsFromDeckResponse(drawnCards)

	response, err := json.Marshal(&drawCardsFromDeckResponse)
	if err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func getDrawnCards(deck models.Deck, numOfDrawnCards int) []models.Card {
	return deck.Cards[len(deck.Cards)-numOfDrawnCards:]
}

func buildDrawCardsFromDeckResponse(drawnCards []models.Card) responses.DrawCardsFromDeckResponse {
	return responses.DrawCardsFromDeckResponse{
		Cards: drawnCards,
	}
}
