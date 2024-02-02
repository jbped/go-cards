package main

import (
	"os"
	"reflect"
	"testing"
)

func TestBuildDeck(t *testing.T) {
	d := buildDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	firstCard := d[0]
	lastCard := d[len(d) - 1]

	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card in deck to be \"Ace of Spades\", but got \"%v\"", firstCard)
	}

	if lastCard != "King of Clubs" {
		t.Errorf("Expected last card in deck to be \"King of Clubs\", but got \"%v\"", lastCard)
	}
}

func TestShuffle(t *testing.T) {
	defaultDeck := deck{"Card 1", "Card 2", "Card 3", "Card 4"}
	testDeck := make(deck, len(defaultDeck))
	copy(testDeck, defaultDeck)

	// Shuffle the deck
	testDeck.shuffle()

	// Suspect to fail if random results in perfect order
	if reflect.DeepEqual(testDeck, defaultDeck) {
		t.Errorf("Expected the deck to be shuffled")
	}
}

func TestDrawCards(t *testing.T) {
	d := deck{"Card 1", "Card 2", "Card 3", "Card 4"}

	// Draw 2 cards from the deck
	hand, remainingDeck := drawCards(d, 2)

	if len(hand) != 2 {
		t.Errorf("Expected hand length of 2, but got %d", len(hand))
	}

	if len(remainingDeck) != 2 {
		t.Errorf("Expected remaining deck length of 2, but got \"%d\"", len(remainingDeck))
	}

	// Add more assertions to validate the contents of the hand
	// and remaining deck based on your implementation
}


func TestToString(t *testing.T) {
	// Create a deck
	d := deck{"Card 1", "Card 2", "Card 3", "Card 4"}

	// Convert the deck to a string
	deckString := d.toString()

	assert := "Card 1,Card 2,Card 3,Card 4"

	if deckString != assert {
		t.Errorf("Expected slice to be converted to string, \"%v\", but got \"%v\"", assert, deckString)
	}
}

func TestSaveDeckAndLoadDeck(t *testing.T) {
	deckName := "_decktest"
	// Remove any lingering test decks
	os.Remove(deckName)
	
	// Create a deck
	d := deck{"Card 1", "Card 2", "Card 3", "Card 4"}
	
	// Save the deck to a file
	d.saveDeck(deckName)
	
	loadedDeck := loadDeck(deckName)
	
	if !reflect.DeepEqual(loadedDeck, d) {
		t.Errorf("Expected saved deck to be created and loaded, but the loadedDeck did not match the deck saved")
	}
	os.Remove(deckName)
}