package round 

import (
	"fmt"
	"../deck"
	// "math"
	// s "strings"
	// "errors"
	// "strconv"
)

type Player struct {
	Name string
	Hand []deck.Card
	Finished bool
	ID string
	// Hand string
}

func (p *Player) RemoveFromHand (indices []int) {
	var newHand []deck.Card 

	for _, index := range indices {
		newHand = p.Hand[:index]
		newHand = append(newHand, p.Hand[index+1:]...)
		p.Hand = newHand
	}
	return 
}

func (p *Player) CardsInHand () int {
	return len((p).Hand)
}

func (p *Player) ShowHand () {
	for i, card := range p.Hand {
		fmt.Print(" | ")
		fmt.Print(i, "-", card.ToStr())
	}
	fmt.Println(" ||")

	return
}

func ShowHand (cards []deck.Card) {
	for i, card := range cards {
		fmt.Print(" | ")
		fmt.Print(i, "-", card.ToStr())
	}
	fmt.Println(" ||")

	return
}