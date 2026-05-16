package main

func main() {
    cards := newDeck()

	hand, remainingcards := deal(cards, 13)
	hand.print()
	remainingcards.print()
}

