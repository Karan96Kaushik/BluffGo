package deck

import (
	_"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit string
	Value int
	Face string
}

type Deck struct {
	Cards []Card
}

func (d *Deck) Initialize () {

	suits := [4]string{"♥", "⬩", "♠", "♣"}
	faces   := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	values  := [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for _, suit := range suits {
		for index, value := range values {
			card := Card{ Suit: suit, Value: value, Face: faces[index] }
			d.Cards = append(d.Cards, card)
		}
	}

	return
}

func (d *Deck) Shuffle () {
	var cards []Card

	totalCards := len(d.Cards)

	for i:=0; i<totalCards; i++ {
		rand.Seed(time.Now().UnixNano())
		randPos := rand.Intn(len(d.Cards))
		// Add random card to new Deck
		cards = append(cards, d.Cards[randPos])
		// Remove card from Deck
		d.Cards[randPos] = d.Cards[len(d.Cards)-1]
		d.Cards = d.Cards[:len(d.Cards)-1]

	}

	d.Cards = cards

	return
}

func (c *Card) ToStr () string {
	return c.Suit + c.Face
}