package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Crete a new type deck
	<<< a slice of string
*/

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{
		"spades",
		"diamonds",
		"hearts",
		"club"}
	dimensions := []string{
		"Ace", "Jack", "Qween", "King",
		"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "ten"}
	for _, suit := range suits {
		for _, value := range dimensions {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), " , ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromFile(d deck, file string) deck {
	dby, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(dby), " , ")
	MyDeck := deck(ss)
	return MyDeck
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
