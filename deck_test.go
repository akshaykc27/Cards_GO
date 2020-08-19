//To make a test,create a new file ending in _test.go

package main

import "testing"

func TestNewDeck(t *testing.T) { // t is testHandler
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length 12 but got %v", len(d))
	}
}
