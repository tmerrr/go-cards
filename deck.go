package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of Deck, will be Slice of strings
type deck []string

func newDeck() deck {
	d := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			d = append(d, cardValue+" of "+cardSuit)
		}
	}

	return d
}

// receiver function
// d is similar to 'this' or 'self' but this convention not used in go
// convention is to normally use 1 or 2 letter name for variable
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function takes deck and integer
// returns 2 instances of deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile("./db/"+filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile("./db/" + filename)
	if err != nil {
		fmt.Println("ERROR READING DECK FROM FILE:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// create random seed source value
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	r := rand.New(source)

	for i := range d {
		// swap current element with random element
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
