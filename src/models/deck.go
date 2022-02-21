package models

type Deck struct {
	UUID      string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

func (d Deck) Suits() [4]string {
	return [4]string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
}

func (d Deck) Values() [13]string {
	return [13]string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
}
