package main

func main() {
	// cards := newDeck()

	// hand, cards := deal(cards, 2)
	// hand.print()
	// cards.print()
	// cards.saveToFile("test.txt")
	cards := newDeckFromFile("test.txt")
	cards.shuffle()
	cards.print()
}
