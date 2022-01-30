package main

import (
	"fmt"
	// "deck"
	"game"
	tcp "tcpServer"
)

func main () {

	p1 := game.Player{	Name: "Bayonet"	, ID: "789"}
	p2 := game.Player{	Name: "Baron" , ID: "345"}
	p3 := game.Player{	Name: "Baroness" , ID: "123"}

	players := []*game.Player{&p1,&p2,&p3}

	go game.StartGame(&players)

	message := make(chan tcp.Message) 
	go tcp.Initialize(message)
	for {
	    msg := <-message
	    if len(msg.Data) > 0 {
		    fmt.Println(msg)
		    // game.ReceivedMove()
	    }
	}
	return

}