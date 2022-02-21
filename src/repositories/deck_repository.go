package repositories

import (
	"deck-api/src/models"
)

type DeckRepository struct{}

var allDecks = []models.Deck{
	{
		UUID: "1",
	    Shuffled: false,
	    Remaining: 1,
	    Cards: []models.Card{
	    	{
	    		Value: "ACE",
	    		Suit: "SPADES",
	    		Code: "AS",
	    	},
	    },
	},
}

func (dr DeckRepository) GetDecks() []models.Deck {
	return allDecks
}

func (dr DeckRepository) SetDecks(decks []models.Deck) {
	allDecks = decks
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Deck struct {
	UUID      string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}