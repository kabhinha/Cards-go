package main

func main() {
	cards := newDeck() // deck.go
	hand, _ := deal(cards, 5)
	hand.shuffle()
	hand.print()
}
