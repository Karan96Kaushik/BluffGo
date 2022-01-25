package game 

import (
	"bufio"
    "fmt"
    "os"
	"deck"
	// "math"
	s "strings"
	"errors"
	"strconv"
)

type Move struct {
	MoveType string
	CardIndices []int
}

func GetClaim () string {
	var claim string

	fmt.Println("Enter Your claim: ")

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
	text = s.Split(text, "\n")[0]
	claim = s.Split(text, " ")[0]

	e := ValidateClaim(claim)
	if e != nil {
		fmt.Println(e)
		return GetClaim()
	}

	return claim
}

func ValidateClaim (claim string) error {
	if len(claim) < 0 {
		return errors.New("Empty claim!")
	}

	for _,f := range deck.Faces {
		if f == claim {
			return nil
		}
	}
	return errors.New("Unrecognized face card!")
}

func GetMove () Move {
	play := []string{}

	fmt.Println("Enter Your Move: ")

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
	text = s.Split(text, "\n")[0]
	play = s.Split(text, " ")

	e := ValidateMove(play)
	if e != nil {
		fmt.Println(e)
		return GetMove()
	}
	move := Move{ MoveType: play[0], CardIndices: []int{} }

	if move.MoveType == "play" {
		for _, val := range play[1:] {
			i, _ := strconv.Atoi(val)
			move.CardIndices = append(move.CardIndices, i)
		}
	}
	return move
}

func ValidateMove (move []string) error {

	possibleTypes := []string{"play", "call", "pass"}
	
	correctType := false
	for _, mType := range possibleTypes {
		if move[0] == mType {
			correctType = true
			break
		}
	}

	if !correctType {
		return errors.New("Invalid move type")
	}

	if move[0] == "play" {
		cardIndex, err := strconv.Atoi(move[1])
		if err != nil {
			return err
		}

		if cardIndex > 52 {
			return errors.New("Invalid card index")
		}
	}

	return nil 
}