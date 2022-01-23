package round 

import (
	"bufio"
    "fmt"
    "os"
	// "../deck"
	// "math"
	s "strings"
	"errors"
	"strconv"
)

type Move struct {
	MoveType string
	CardIndices []int
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