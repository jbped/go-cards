package main

import(
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // Preferred way to declare a variable
	card := newCard()
	fmt.Println(card)
}
func newCard() string {
	return "Five of Diamonds"
}