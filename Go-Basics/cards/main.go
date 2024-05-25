package main

import "fmt"

// Function will create a deck of card, shuffle it and give you hand of five card and also print the remaining cards of the deck
func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	cards.shuffle()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("********************************")
	remainingDeck.print()
	// cards = newDeckFromFile("my_cards")
	// cards.print()
}
