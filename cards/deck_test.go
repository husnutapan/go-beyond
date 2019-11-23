package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Excepted deck length of 16, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("testdeck")

	deck := newDeck()
	deck.saveToFile("testdeck")

	loadDeck := newDeckFromFile("testdeck")

	if len(loadDeck) != 16 {
		t.Errorf("Expecting 16 cards in deck, but got %v", len(loadDeck))
	}
	os.Remove("testdeck")
}
