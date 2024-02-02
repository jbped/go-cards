package main

import(
	"fmt"
)

func main() {
	cards := []string{"Ace of Spades", newCard()}
	fmt.Println(cards)
	// iterate over the slice, and print
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
func newCard() string {
	return "Five of Diamonds"
}