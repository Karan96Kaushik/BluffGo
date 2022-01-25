package game 

import (
	"fmt"
	"deck"
	"math"
	// _ "strings"
	// _ "errors"
)

type Round struct {
	CurrentPlayer *Player
	LastPlayer *Player
	Floor []deck.Card
	Claim string
	Ended bool
	Players *[]*Player
	LastMove []deck.Card
	ContinousPasses int
}

func StartGame (players *[]*Player) {

	totalPlayers := len(*players)
	fmt.Println("Starting game...")
	fmt.Println("Num of players: ", totalPlayers)

	var d deck.Deck

	d.Initialize()
	d.Shuffle()

	hands := d.Deal(totalPlayers, int(math.Round(52/float64(totalPlayers))))
	// fmt.Println(hands)

	for index, cards := range hands {
		(*players)[index].AddToHand(cards)
		fmt.Println((*players)[index].Name, (*players)[index].CardsInHand())
		(*players)[index].ShowHand()
	}

	StartRound(players)
	
	return

}

func StartRound (players *[]*Player) {
	var r Round

	r.CurrentPlayer = (*players)[0]
	r.Ended = false
	r.Players = players

	for !r.Ended {

		// fmt.Println(r.CurrentPlayer, r.LastPlayer)
		fmt.Print(r.CurrentPlayer.Name, ": ")
		r.CurrentPlayer.ShowHand()
		if len(r.Floor) == 0 {
			r.Claim = GetClaim()
		}
		move := GetMove()
		r.ProcessMove(move)
		// fmt.Println(r.CurrentPlayer, r.LastPlayer)
		// r.LastPlayer = r.CurrentPlayer
		// fmt.Println("B", r.CurrentPlayer, r.LastPlayer)
		// r.LastPlayer = nil
		// fmt.Println("YN", r.CurrentPlayer, r.LastPlayer)
	}

	// fmt.Print("TCURR")
	// turn.CurrentPlayer.ShowHand()
	// fmt.Print("Player")
	// (*players)[0].ShowHand()

	return
}

func (r *Round) ProcessMove (move Move) {

	if move.MoveType == "play" {
		r.LastMove = []deck.Card{}
		for _, index := range move.CardIndices {
			r.Floor = append(r.Floor, r.CurrentPlayer.Hand[index])
			r.LastMove = append(r.LastMove, r.CurrentPlayer.Hand[index])
		}
		r.CurrentPlayer.RemoveFromHand(move.CardIndices)
		r.NextPlayer()
	} else if move.MoveType == "call" {
		r.CallBluff()
	} else if move.MoveType == "pass" {
		r.NextPlayer()
	}
	
	if move.MoveType != "pass" {
		r.ContinousPasses = 0
	}

	fmt.Print("Floor: ")
	ShowHand(r.Floor)
}

func (r *Round) NextPlayer () {
	r.LastPlayer = r.CurrentPlayer
	for i, p := range *r.Players {
		if p.ID == (*r.CurrentPlayer).ID {
			if i >= len(*r.Players)-1 {
				r.CurrentPlayer = (*r.Players)[0]
			} else {
				r.CurrentPlayer = (*r.Players)[i+1]
			}
			break
		}
	}

}

func (r *Round) CallBluff () {
	fmt.Println(r.CurrentPlayer.Name,(*r.LastPlayer).Name)

	for _, card := range r.LastMove {
		// If last move was a bluff
		if card.Face != r.Claim {
			(*r.LastPlayer).AddToHand(r.Floor)
			r.Floor = []deck.Card{}
			fmt.Println("Bluff Caught by", r.CurrentPlayer.Name + "!")
			return
		}
	}
	// Not a bluff
	fmt.Println("Not a Bluff!", "Next Player", (*r.LastPlayer).Name)
	r.CurrentPlayer.AddToHand(r.Floor)
	r.CurrentPlayer = r.LastPlayer
	r.Floor = []deck.Card{}
	return
}