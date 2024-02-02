package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Ace of Hearts")

	cards.print()
}
func newCard() string {
	return "Ace of Diamonds"
}