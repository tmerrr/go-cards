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

// select range from slice
// sliceName[startIndexIncluding : upToNotIncluding]
// fruits := []string{"apple", "banana", "grape", "orange"}
// fruits[0:2] same as fruits[:2] - start from begining
// => ["apple", "banana"]
// fruits[2:] - go to end
// => ["grape", "orange"]

// function takes deck and integer
// returns 2 instances of deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

// could pass in filename, decide later
func (d deck) saveToFile(filename string) error {
	// last argument is permissions
	// 0666 means anyone can read or write to file
	return ioutil.WriteFile("./db/"+filename, []byte(d.toString()), 0666)
}

// Could take filename as arg
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

	// iterate over array
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap current element with random element
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
