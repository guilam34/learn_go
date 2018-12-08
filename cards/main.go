package main

func main() {
	// Slice
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
