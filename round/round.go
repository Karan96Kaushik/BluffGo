package round 

import (
	"fmt"
	"../deck"
	"math"
	// _ "strings"
	// _ "errors"
)

type Turn struct {
	CurrentPlayer *Player
	Floor []deck.Card
	Ended bool
	Players *[]Player
	LastMove []deck.Card
}

func StartRound (players *[]Player) {

	totalPlayers := len(*players)
	fmt.Println("Starting round...")
	fmt.Println("Num of players: ", totalPlayers)

	var d deck.Deck

	d.Initialize()
	d.Shuffle()

	hands := d.Deal(totalPlayers, int(math.Round(52/float64(totalPlayers))))
	// fmt.Println(hands)

	for index, cards := range hands {
		(*players)[index].Hand = cards
		fmt.Println((*players)[index].Name, (*players)[index].CardsInHand())
		(*players)[index].ShowHand()
	}
	
	var turn Turn
	turn.CurrentPlayer = &(*players)[0]
	turn.Ended = false
	turn.Players = players

	for !turn.Ended {

		fmt.Print(turn.CurrentPlayer.Name, ": ")
		turn.CurrentPlayer.ShowHand()
		move := GetMove()
		turn.ProcessMove(move)
		turn.NextPlayer()
	}

	// fmt.Print("TCURR")
	// turn.CurrentPlayer.ShowHand()
	// fmt.Print("Player")
	// (*players)[0].ShowHand()

	return

}

func (t *Turn) ProcessMove (move Move) {

	if move.MoveType == "claim" {
		t.LastMove = []deck.Card{}
		for _, index := range move.CardIndices {
			t.Floor = append(t.Floor, t.CurrentPlayer.Hand[index])
			t.LastMove = append(t.LastMove, t.CurrentPlayer.Hand[index])
		}
		t.CurrentPlayer.RemoveFromHand(move.CardIndices)
	} else if move.MoveType == "play" {
		t.LastMove = []deck.Card{}
		for _, index := range move.CardIndices {
			t.Floor = append(t.Floor, t.CurrentPlayer.Hand[index])
			t.LastMove = append(t.LastMove, t.CurrentPlayer.Hand[index])
		}
		t.CurrentPlayer.RemoveFromHand(move.CardIndices)
	} else if move.MoveType == "call" {

	}
	
	fmt.Print("Floor: ")
	ShowHand(t.Floor)
}

func (t *Turn) NextPlayer () {

	for i, p := range *t.Players {
		if p.ID == (*t.CurrentPlayer).ID {
			if i >= len(*t.Players)-1 {
				t.CurrentPlayer = &(*t.Players)[0]
			} else {
				t.CurrentPlayer = &(*t.Players)[i+1]
			}
			break
		}
	}

}