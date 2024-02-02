package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

// Create a deck of cards
// Return a deck (slice of strings)
func buildDeck() deck {
	cardDeck := deck{}
	suites := []string{
		"Spades", 
		"Diamonds", 
		"Hearts", 
		"Clubs",
	}
	values := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suite := range suites {
		for _, value := range values {
			cardDeck = append(cardDeck, value + " of " + suite)
		}
	}

	return cardDeck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
	fmt.Println("---------------")
}

func drawCards(d deck, amt int) (deck, deck) {
	cardsDrawn := d[:amt]
	deckRemaining := d[amt:]

	return cardsDrawn, deckRemaining
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeck(deckName string) error {
	return os.WriteFile(deckName + ".txt", []byte(d.toString()), 0666)
}
