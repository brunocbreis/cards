package main

import (
	"fmt"
	"strconv"
)

// Define a Card
type Card struct {
	Value int    `json:"value"`
	Suit  string `json:"suit"`
}

var cardValues = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
var cardSuits = [4]string{"Hearts", "Spades", "Diamonds", "Clubs"}

func (c Card) name() string {
	// Returns a human-readable name for a card.

	letterCards := map[int]string{1: "Ace", 11: "Jack", 12: "Queen", 13: "King"}

	var name string

	if value, is_letter := letterCards[c.Value]; is_letter {
		name = value
	} else {
		// convert integer card value to string
		name = strconv.Itoa(c.Value)
	}

	return fmt.Sprintf("%s of %s", name, c.Suit)
}
