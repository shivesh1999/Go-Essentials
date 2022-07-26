package main

func main() {
	// we can intialize a variable outside the function but cannor declare it
	// primitive data structures - array and slice
	// array is fixed length, slice is array that can grow or shrink
	// both array or slice need a datatype to be defined
	//var card string = "Ace of spades"
	//cards := []string{"Ace of diamonds", newCard()}
	// here a new slice is created when we append the value so we have to assign t to a variable
	//cards = append(cards, "sex of spades")
	// i is the index of the element in array
	// card is the current card we are iterating
	// range cards is the slice that we are looping over
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	//cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

//functions in Go
// func newCard() string {
// 	return "Five of Diamonds"
// }

//struct
// data structure
// colllection of properties related to each other
