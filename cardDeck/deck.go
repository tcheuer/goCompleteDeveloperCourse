package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

/*
	The deck type is used to simulate
	a deck of cards, the leftmost string
	being the top card and the rightmost
	being the bottom card.
*/
type deck []string

/*
	Prints the contents of the calling deck
	to the console.
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
	Generates a new 'standard' deck of playing cards
	(52 cards, 4 suits of 13 each)
*/
func newDeck() deck {
	cards := deck{}
	suits := []string{
		"Spades",
		"Clubs",
		"Hearts",
		"Diamonds"}
	values := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}

	for _, s := range suits {
		for _, v := range values {
			cards = append(cards, v+" of "+s)
		}
	}

	return cards
}

/*
	Creates a new "hand" (deck object) by
	splitting off the first drawNumber cards
	into a separate slice (which is returned),
	and defining an updated deck using the remaining
	cards.

	Returns hand, deck, error
*/
func deal(d deck, drawNumber int) (hand deck, newDeck deck, err error) {
	if drawNumber <= 0 {
		return nil, d, errors.New("cannot deal 0 or less cards")
	}

	hand = d[0:drawNumber]
	newDeck = d[drawNumber:]

	return hand, newDeck, nil
}

/*
	Saves the passed deck d to a file fileName in bytes.
*/
func (d deck) writeToFile(fileName string) error {

	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

	return err

}

/*
	Converts the calling deck to a single string.
*/
func (d deck) toString() string {
	return strings.Join([]string(d), "::")
}

/*
	Splits a string along the :: character and returns
	it as a deck object.
*/
func deckFromString(dString string) deck {
	return deck(strings.Split(dString, "::"))
}

/*
	Reads a deck from a file, and returns it as a new deck object

	If reading fails, returns nil and an error message.
*/
func newDeckFromFile(fileName string) (deck, error) {
	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	d := deckFromString(string(byteSlice))

	return d, nil
}

/*
	Shuffles the values in the calling deck using
	Intn from the math/rand library.
*/
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randIndex := r.Intn(len(d) - 1)

		//Swaps two elemenets without losing either
		d[i], d[randIndex] = d[randIndex], d[i]
	}
}
