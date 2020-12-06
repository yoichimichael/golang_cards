// will an executable-type package
package main

import "fmt"

// unshowable, but VSCode added this automatically when saving after writing fmt.Println(card)

// variabel declaration without value assignment can occur outside of a func
// var card string

// every executable type package needs to contain a func called main
func main() {
	// long form variable declaration and value assignment
	// var card string = "Ace of Spades"

	// value reassignment doesn't use : (colon)
	// card = "Five of Diamonds"

	// alternative variabel declaration syntax
	// Go infers data type set to variable by the value assigned to it with :=
	cards := []string{"Ace of Diamonds", newCard()}

	//necessary to store result of append, often in variable holding the slice
	cards = append(cards, "Six of Spades")

	// slice iteration
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	for index, card := range cards {
		fmt.Println(card)
	}

	// fmt.Println(cards)
}

// return type needs to be declared
// without 'string' after 'newCard()' the function is telling the compiler that it doesn't want to return anything (in spite of the return statement)
func newCard() string {
	return "Five of Diamonds"
}
