package main

import (
	"os"
	"testing"
)

var deckLen = 28

func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != deckLen {
		t.Errorf("Expected length of deck to be %v but got %v",deckLen, len(d))
	}
}
func TestFirstAndLastCard(t *testing.T){
	d := newDeck()
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}
	if d[len(d) - 1] != "Seven of Clubs"{
		t.Errorf("Expected first card to be Seven of Clubs but got %v", d[0])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	fileName := "_decktesting"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)
	loadedDeck := newDeckFromFile(fileName)
	if len(loadedDeck) != deckLen{
		t.Errorf("Expected length of loaded deck to be %v but got %v",deckLen, len(loadedDeck))
	} 
	os.Remove(fileName)
}