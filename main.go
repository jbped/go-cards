package main

import(
	"fmt"
)

func main() {
	cards := []string{"Ace of Spades", newCard()}
	fmt.Println(cards)
	// iterate over the slice, and print
	for _, card := range cards {
		fmt.Println(card)
	}
}
func newCard() string {
	return "Five of Diamonds"
}