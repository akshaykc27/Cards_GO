package main

func main() {
	cards := newDeck()
	// cards.print()

	// hand, remaining := deal(cards, 7)
	// fmt.Println(hand)
	// fmt.Println(remaining)

	cards.saveToFile("CardGame")
}
