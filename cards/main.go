package main

import "fmt"

func main()  {

	cards := newDeck()


	//cards.print()

	hand, remainingDeck := cards.deal(5)
	hand.print()
	fmt.Println("---")
	remainingDeck.print()
	fmt.Println("===")

	newHand, remDeck := hand.deal(2)
  	newHand.print()
	fmt.Println("---")
  	remDeck.print()


}




func newCard() string  {
	return "Five of Diamonds"
}
