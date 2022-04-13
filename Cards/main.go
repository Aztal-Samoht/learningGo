package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	cards := deck{newCard(), newCard(), "ace of spades"	}
	fmt.Println(cards)
	cards = append(cards, "six of spades")
	cards.addQ()
	cards.print()
}
	

func newCard() string {
	return "five of diamonds"
}