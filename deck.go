package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

// receiver function
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//receiver function
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// assigning bs (byte slice) and err (error) to return values of ReadFile()
	bs, err := ioutil.ReadFile(filename)
	// below is very common pattern for error handling
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// notice type conversion: bs (byte slice) to string, and s (slice of strings) to deck type
	s := strings.Split(string(bs), ",")
	return deck(s)
}
