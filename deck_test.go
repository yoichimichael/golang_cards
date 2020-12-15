package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// tests for length of deck
	// note string interpolation with %v"
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// tests for correct first card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// tests for correct last card
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
