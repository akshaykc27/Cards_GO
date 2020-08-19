package main

import "fmt"

//Create a new type  "deck"
//which is slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suite := range cardSuites {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suite)
		}
	}

	return cards
}

//receiver (d deck) any variable of type deck will have the access of the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
