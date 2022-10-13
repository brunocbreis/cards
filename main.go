package main

import "fmt"

func main() {
	deck := newDeck()
	deck.shuffle()

	var hand Deck
	var flop Deck

	hand = deck.deal(2)

	flop = deck.deal(5)

	fmt.Printf("Your hand is:\t%s\n", hand.toString())
	fmt.Printf("The flop is:\t%s\n", flop.toString())
	fmt.Printf("The remaining deck has %d cards.\n", len(deck))
}
