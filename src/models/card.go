package models

import (
	"fmt"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func (c Card) GetCode() string {
	return fmt.Sprintf("%s%s", c.Value[0:1], c.Suit[0:1])
}
