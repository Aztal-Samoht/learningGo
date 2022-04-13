package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	cards := []string{newCard(), newCard(), "ace of spades"	}
	fmt.Println(cards)
	cards = append(cards, "six of spades")
	for i, aCard := range cards{
		fmt.Println(i)
		fmt.Println(aCard)
	}
}
	

func newCard() string {
	return "five of diamonds"
}