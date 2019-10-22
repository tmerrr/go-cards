package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFile("shuffledDeck")
	hand, remainingCards := deal(cards, 7)
	cards = remainingCards
	cards.saveToFile("remainingCards")

	hand.print()
	cards.print()

	hand.saveToFile("hand")
	loadedDeck := newDeckFromFile("remainingCards")
	loadedDeck.print()
}
