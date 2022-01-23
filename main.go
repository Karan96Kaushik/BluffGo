package main

import (
	_"fmt"
	_"./deck"
	"./round"
)



func main () {

	p1 := round.Player{	Name: "Bayonet"	, ID: "789"}
	p2 := round.Player{	Name: "Baron" , ID: "345"}
	p3 := round.Player{	Name: "Baroness" , ID: "123"}

	players := []round.Player{p1,p2,p3}
	round.StartRound(&players)
	return
}