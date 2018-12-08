package main

import (
	"fmt"
)

func main() {
	// type declaration isn't needed since variable assignment infers type
	// var card = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}
