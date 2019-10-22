package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// check length of slice
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	// check value of first element (Ace of Spades)
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be Ace of Spades, but got %v", d[0])
	}
	// check value of last element (King of Clubs)
	if d[51] != "King of Clubs" {
		t.Errorf("Expected last element to be King of Clubs, but got %v", d[51])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove("./db/" + filename)
	fmt.Println("Deleted _decktesting file from db")

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected Loaded Deck to have length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("./db/" + filename)
}
