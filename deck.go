package main

import (
	"fmt"
	"strings"
)

// create/declare a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// underscores represent unused variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// "function 'deal' takes arguments 'd' of type deck and 'handsize' of type int and returns two values of type deck"
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
