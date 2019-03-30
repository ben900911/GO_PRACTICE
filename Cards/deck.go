package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

/*
 cards : copy of the deck we're owrking with is available in the function as a var called cards ()reference)
this, self in OOP

 deck : every var of type 'deck' can  call this function on itself

*/
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (cards deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[index], cards[newPosition] = cards[newPosition], cards[index]
	}
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(cards.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteslice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteslice), ",")
	return deck(s)
}
