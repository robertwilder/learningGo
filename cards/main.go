package main

import "fmt"

var deckSize int

func main() {

	card := newCard()

	fmt.Println(card)


}


func newCard() string {
	return "Five of Diamonds"
}
	// var card string = "Ace of Spades"
	// Two means of using variables 
	// only use := when you are assigning a new variable 
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// fmt.Println(card)


	// You can Initialize the variable later
	// var deckSize int
//   deckSize = 52
//   fmt.Println(deckSize)

// variable can be intialized out of the the func main()
// deckSize = 50
// fmt.Println(deckSize)

// Go is a statically typed language

// Bool 
// string 
// int 
// float64 - number with a decimal after it. 10.000001
