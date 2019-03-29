package main

import (
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "strings"
  "time"
)

//Create a new type of "deck"
//which is a slice of strings

type deck []string

func newDeck () deck   {
  cards := deck{}

  cardSuits := []string {"Spades", "Hearts", "Diamonds", "Clubs"}
  cardValues := []string {
    "Ace",  "Two", "Three", "Four",
    "Five", "Six", "Seven", "Eight", "Nine", "Ten",
    "King", "Queen", "Joker",
  }

  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards  = append(cards, (value+" of "+suit))
    }
  }

  return cards
  }

func (d deck) print(){
  for i, card := range d {
    fmt.Println((i+1), card  )
  }
}

//Takes in a deck of cards as a "Receiver" and and handsize as an "Argument"
func deal(d deck, handSize int) (deck, deck){
  return d[:handSize], d[handSize:]
}

//Takes a deck of cards as a "Receiver"
func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}

func toByte(d string) []byte {
  return []byte(d)
}

//Save Deck of cards to hardDrive
//0666 is a permision that says "anyone can read or write this file"
func (d deck) saveToFile (filename string) error  {
  return ioutil.WriteFile(filename, [] byte(d.toString()), 0666)
}

//Retrieve deck of cards from a file in hardDrive
func newDeckFromFile(filename string) deck {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    //Option #1 - log the error and make a call to newDeck.
    //Option #2 - log the error and entirely quit the program.
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  s := strings.Split(string(bs), ",")

  return deck(s)
}

func (d deck) shuffle() {
    source := rand.NewSource(time.Now().UnixNano())
    r :=  rand.New(source)

   for i := range d{
     newPosition := r.Intn(len(d)-1)
     d[i], d[newPosition] = d [newPosition], d[i]
   }

}