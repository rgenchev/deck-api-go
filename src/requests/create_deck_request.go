package requests

import (
	"encoding/json"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strings"

	"deck-api/src/models"
	"deck-api/src/repositories"
	"deck-api/src/responses"
)

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var deckRepository repositories.DeckRepository
	decks := deckRepository.GetDecks()

	selectedCardCodesParam := r.URL.Query().Get("cards")
	var newDeck models.Deck
	var selectedCardCodes []string

	err := json.NewDecoder(r.Body).Decode(&newDeck)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if len(selectedCardCodesParam) > 0 {
		selectedCardCodes = strings.Split(selectedCardCodesParam, ",")
	}

	newDeck.UUID = uuid.New().String()
	newDeck.Cards = make([]models.Card, 0)

	newDeck.Cards = buildCards(newDeck, selectedCardCodes)

	if newDeck.Shuffled {
		newDeck = shuffleDeck(newDeck)
	}

	newDeck.Remaining = len(newDeck.Cards)

	decks = append(decks, newDeck)
	deckRepository.SetDecks(decks)

	newDeckResponse := buildNewDeckResponse(newDeck)

	response, err := json.Marshal(&newDeckResponse)
	if err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func buildCards(deck models.Deck, selectedCardCodes []string) []models.Card {
	for _, suit := range deck.Suits() {
		for _, value := range deck.Values() {
			newCard := models.Card{
				Value: value,
				Suit:  suit,
			}

			newCard.Code = newCard.GetCode()

			if len(selectedCardCodes) > 0 {
				deck.Cards = appendSelectedCardToDeck(deck, newCard, selectedCardCodes)
			} else {
				deck.Cards = append(deck.Cards, newCard)
			}
		}
	}

	return deck.Cards
}

func shuffleDeck(deck models.Deck) models.Deck {
	for i := range deck.Cards {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}

	return deck
}

func buildNewDeckResponse(deck models.Deck) responses.NewDeckResponse {
	return responses.NewDeckResponse{
		UUID:      deck.UUID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}
}

func appendSelectedCardToDeck(deck models.Deck, selectedCard models.Card, selectedCardCodes []string) []models.Card {
	for _, cardCode := range selectedCardCodes {
		if selectedCard.Code == cardCode {
			deck.Cards = append(deck.Cards, selectedCard)
		}
	}

	return deck.Cards
}
