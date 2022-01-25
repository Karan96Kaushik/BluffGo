package deck

import (
	"fmt"
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

var Suits [4]string 	= [4]string{"♥", "⬩", "♠", "♣"}
var Faces [13]string 	= [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var Values [13]int 		= [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

func (d *Deck) Initialize () {
	for _, suit := range Suits {
		for index, value := range Values {
			card := Card{ Suit: suit, Value: value, Face: Faces[index] }
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

func (d *Deck) Deal (numOfHands int, cardsPerHand int) [][]Card {
	fmt.Println("Dealing cards...")
	var hands [][]Card
	
	totalCards := numOfHands * cardsPerHand
	if totalCards > len(d.Cards) {
		totalCards = len(d.Cards)
	}

	// Initialize each hand
	for i:=0; i<numOfHands; i++ {
		var cards []Card
		hands = append(hands, cards)
	}

	// Deal cards for each hand
	for i := 0; i < (totalCards+1); i++ {
		// fmt.Println(i)
		hands[i % numOfHands] = append(hands[i % numOfHands], d.Cards[i])
	}

	d.Cards = d.Cards[totalCards - 1:]
	return hands
}

func (c *Card) ToStr () string {
	return c.Suit + c.Face
}