//To make a test,create a new file ending in _test.go

package main

import "testing"

func TestNewDeck(t *testing.T) { // t is testHandler
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length 12 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v ", d[0])
	}

	if d[len(d)-1] != "Three of Clubs" {
		t.Errorf("Expected last card to be Three of Clubs, but got %v ", d[len(d)-1])
	}
}
