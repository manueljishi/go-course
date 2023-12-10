package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new deck type

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}
func newDeckFromFile(fileName string) deck {
	res, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	//Res has the deck encoded to a byte slice
	deckStr := string(res)
	return strings.Split(deckStr, ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
