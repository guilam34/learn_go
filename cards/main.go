package main

import (
	"fmt"
)

func main() {
	// type declaration isn't needed since variable assignment infers type
	// var card = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
