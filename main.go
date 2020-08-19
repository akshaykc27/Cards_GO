package main

import "fmt"

func main() {
	cards := newDeckFromFile("CardGame")
	// cards.print()

	// hand, remaining := deal(cards, 7)
	// fmt.Println(hand)
	// fmt.Println(remaining)
	fmt.Println(cards)
}
