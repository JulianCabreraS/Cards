package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deckExample := newDeck()
	if len(deckExample) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deckExample))
	}

	if deckExample[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deckExample[0])
	}
}

func TestElementSlice(t *testing.T) {
	deckExample := newDeck()

	if deckExample[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deckExample[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
