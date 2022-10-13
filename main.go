package main

func main() {
	// deck := newDeck()

	// var hand Deck

	// hand, deck = deck.deal(5)

	// fmt.Println("Your hand is:")
	// hand.print()
	// fmt.Printf("\nThere are %d cards left in the deck:\n", len(deck))
	// fmt.Println(deck.toString())

	// hand.toJSON("myHand.json")

	importedHand, _ := newDeckFromJSON("myHand.json")
	importedHand.print()

}
