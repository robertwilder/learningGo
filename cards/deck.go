package main

import "fmt"

//  create a new type of 'deck'
// which is a slice of strings


// deck is kind of like a sub class of the type string
type deck[]string


func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs" }
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	// if you don't need to use a variable but it needs to be there, use an '_'

	return cards
} 

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}




// func (d deck) print()
// (d deck) is the receiver 
// the receiver sets up variables on methods we create
// convention of a 1 or two letter reference to the deck
// receiver argument 



// Byte Slice, think in your head string. byte Slice = []byte
// go language expects a byte slice 

//  go Type conversion. - []byte("Hi There!")