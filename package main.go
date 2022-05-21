package main

func main() {
	/*
		var card = newCard()
		fmt.Println(card)
	*/
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("my_cards")

	cards2 := newDeckFromFile("my_cards")
	cards2.shuffle()
	cards2.print()
}

/*func newCard() string {
	return "Five of Diamonds"
}
*/
