package main

func main() {
	cards := buildDeck()
	cards.print()
	cards.saveDeck("new_deck")

	// hand, cards := drawCards(cards, 5)
	// hand.print()
	// cards.print()
}