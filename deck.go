package main

import (
	"fmt"
	"math/rand"
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

func (d deck) shuffle() {
	for i := range d {
		randomIndex := rand.Intn(len(d) - 1)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}

func drawCards(d deck, amt int) (deck, deck) {
	cardsDrawn := d[:amt]
	deckRemaining := d[amt:]

	return cardsDrawn, deckRemaining
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
	fmt.Println("---------------")
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeck(deckName string) error {
	return os.WriteFile("decks/" + deckName + ".txt", []byte(d.toString()), 0666)
}

func loadDeck(deckName string) deck {
	if deckName == "" {
		fmt.Println("Error: Missing deckName")
		os.Exit(1)
	}
	byteDeck, err := os.ReadFile("decks/" + deckName + ".txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	loadedDeck := strings.Split(string(byteDeck), ",")
	return deck(loadedDeck)
}