package main

import "fmt"

// Function will create a deck of card, shuffle it and give you hand of five card and also print the remaining cards of the deck
func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	cards.shuffle()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("********************************")
	fmt.Println("Your hand is: ")
	fmt.Println("********************************")
	hand.print()
	fmt.Println()
	fmt.Println("********************************")
	fmt.Println("Remaining cards: ")
	fmt.Println("********************************")
	remainingDeck.print()
}
