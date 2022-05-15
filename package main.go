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
}

func newCard() string {
	return "Five of Diamonds"
}
