package main

func main() {
	cards := NewDeck() // deck.go
	hand, _ := deal(cards, 5)
	hand.shuffle()
	hand.print()
}
