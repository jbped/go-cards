package main

func main() {
	cards := buildDeck()
	cards.print()

	hand, cards := drawCards(cards, 5)
	hand.print()
	cards.print()
}