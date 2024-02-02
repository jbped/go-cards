package main

func main() {
	cards := buildDeck()
	cards.print()
	// cards.saveDeck("new_deck")
	// cards := loadDeck("new_deck")
	cards.shuffle()

	// hand, cards := drawCards(cards, 5)
	// hand.print()
	cards.print()
}