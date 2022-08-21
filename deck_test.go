package main

import "testing"

func TestnewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 53 {
		t.Errorf("Expected length to be 53 but got %v", len(d))
	}
}
