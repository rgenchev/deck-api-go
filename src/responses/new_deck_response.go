package responses

type NewDeckResponse struct {
	UUID      string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}
