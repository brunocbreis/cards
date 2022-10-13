package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Deck []Card

func newDeck() Deck {
	var d Deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, Card{value, suit})
		}
	}
	return d
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card.name())
	}
}

func (d Deck) deal(handSize int) (Deck, Deck) {
	// Deals a hand and returns the hand and the new deck back

	hand := d[:handSize]
	deck := d[handSize:]

	return hand, deck
}

func (d Deck) toString() string {
	var deckSlice []string
	for _, card := range d {
		deckSlice = append(deckSlice, card.name())
	}

	return strings.Join(deckSlice, ", ")
}

func (d Deck) toJSON(fileName string) error {
	jsonBytes, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(fileName, jsonBytes, 0666)
	return err
}

func newDeckFromJSON(fileName string) (Deck, error) {
	jsonBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var newDeck Deck
	json.Unmarshal(jsonBytes, &newDeck)

	return newDeck, err
}
