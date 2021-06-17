package main

import "fmt"

//  create a new type of 'deck'
// which is a slice of strings


// deck is kind of like a sub class of the type string
type deck[]string

// Receiver  

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}


// func (d deck) print()
// (d deck) is the receiver 
// the receiver sets up variables on methods we create
// convention of a 1 or two letter reference to the deck
// receiver argument 


