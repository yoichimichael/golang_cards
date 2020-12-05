// will an executable-type package
package main

// unshowable, but VSCode added this automatically when saving after writing fmt.Println(card)
import "fmt"

// every executable type package needs to contain a func called main
func main() {
	// value assignment to variable
	// var card string = "Ace of Spades"

	// alternative syntax
	// Go infers data type set to variable by the value assigned to it with :=
	card := "Ace of Spades"
	// value reassignment doesn't use : (colon)
	card = "Five of Diamonds"
	fmt.Println(card)
}
