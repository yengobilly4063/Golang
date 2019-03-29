package main

import (
  "os"
  "testing"
)

func TestNewDeck(t *testing.T)  {
  d := newDeck()


  //Test 1 for Length of Deck
  if len(d) != 52 {
	t.Errorf("Expected Deck length of 52, but got %v", len(d))
  }

  //Test 2 for First card in Deck
  if d[0] != "Ace of Spades" {
	t.Errorf("Expected 'Ace of Spades', but got %v", d[0])
  }

  //Test 3 for last card in Deck
  if d[len(d)-1] != "Joker of Clubs" {
	t.Errorf("Expected 'Joker of Clubs', but got %v", d[len(d)-1])
  }

}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
  os.Remove("_decktesting")

  deck := newDeck()
  deck.saveToFile("_decktesting")

  loadedDeck := newDeckFromFile("_decktesting")

  if len(loadedDeck) != 52 {
	 t.Errorf("Expected length 52, but got %v", len(loadedDeck))
  }

  os.Remove("_decktesting")
}