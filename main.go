package main

func main() {
	cards := newDeck() // deck.go
	hand, _ := deal(cards, 5)
	hand.shuffle()
	hand.saveToFile("abc.txt")
	hand.deckFromFile("abc.txt")
	hand.print()

}
