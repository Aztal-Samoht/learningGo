package main

import "fmt"

//create a new type of deck which is a slice of stings
type deck []string

func (d deck) print(){
	for i, aCard := range d{
		fmt.Println(i, aCard)
	}
}

func (d deck) addQ() deck{
	d = append(d, "queen of hearts")
	return d
	//d.print()
}