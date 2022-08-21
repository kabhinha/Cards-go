package main

func main() {
	cards := newDeck() // deck.go
	cards.shuffle()
	cards.shuffle()
	cards.shuffle()
	hand, _ := deal(cards, 5)
	hand.shuffle()
	hand.saveToFile("abc.txt")
	deckFromFile(hand, "abc.txt")
	hand.print()

}
