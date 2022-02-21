package responses

import (
	"deck-api/src/models"
)

type DrawCardsFromDeckResponse struct {
	Cards []models.Card `json:"cards"`
}
