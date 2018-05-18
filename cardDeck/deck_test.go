package main

import (
	"os"
	"testing"
)

func Test_newDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52 but recieved %d", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be 'Ace of Spades', received %v", cards[0])
	}

	if cards[len(cards)-1] != "King of Diamonds" {
		t.Errorf("Expected final element to be 'King of Diamonds', received %v", cards[len(cards)-1])
	}

}

func Test_writeToFile_newDeckFromFile(t *testing.T) {
	//Remove any previous files leftover from tests
	os.Remove("_decktesting")

	//Check read/write
	d := newDeck()

	err := d.writeToFile("_decktesting")

	if err != nil {
		t.Fatalf("Failed to write deck to file: %v", err)
	}

	loadedDeck, err := newDeckFromFile("_decktesting")

	if err != nil {
		t.Fatalf("Failed to load deck from file: %v", err)
	}

	//Check validity of loaded deck
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 but recieved %d", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be 'Ace of Spades', received %v", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Diamonds" {
		t.Errorf("Expected final element to be 'King of Diamonds', received %v", loadedDeck[len(loadedDeck)-1])
	}

	os.Remove("_decktesting")

}
