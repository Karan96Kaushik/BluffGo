package main

import (
	"fmt"
	"./deck"
)

func main () {
	var d deck.Deck
	d.Initialize()
	d.Shuffle()

	for _, card := range d.Cards {
		fmt.Println(card.ToStr())
	}
}