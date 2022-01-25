package main

import (
	// "fmt"
	// "deck"
	"game"
)

func main () {

	p1 := game.Player{	Name: "Bayonet"	, ID: "789"}
	p2 := game.Player{	Name: "Baron" , ID: "345"}
	p3 := game.Player{	Name: "Baroness" , ID: "123"}

	players := []*game.Player{&p1,&p2,&p3}
	game.StartGame(&players)
	return
}