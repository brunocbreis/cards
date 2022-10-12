package main

import "fmt"

type deck []card

func newDeck() deck {
	var d deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, card{value, suit})
		}
	}
	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card.name())
	}
}
