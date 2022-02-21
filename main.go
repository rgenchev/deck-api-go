package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"deck-api/src/requests"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/decks/{uuid}", requests.GetDeckByUUID).Methods("GET")
	router.HandleFunc("/decks", requests.CreateDeck).Methods("POST")
	router.HandleFunc("/decks/{uuid}/draw/{count}", requests.DrawCardsFromDeck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
