package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var field = createField()
	var turn int = rand.Intn(2)
	for {
		displayField(field)
		makeTurn(&turn, &field)
		if winner := isGameEnd(field); winner != -1 {
			fmt.Printf("Player %d wins! Thanks for playing!\n", winner+1)
			time.Sleep(5 * time.Second)
			break
		}
	}

}

func isGameEnd(field [3][3]string) int {
	if field[0][0] == "X" && field[0][1] == "X" && field[0][2] == "X" ||
		field[1][0] == "X" && field[1][1] == "X" && field[1][2] == "X" ||
		field[2][0] == "X" && field[2][1] == "X" && field[2][2] == "X" ||
		field[0][0] == "X" && field[1][0] == "X" && field[2][0] == "X" ||
		field[0][1] == "X" && field[1][1] == "X" && field[2][1] == "X" ||
		field[0][2] == "X" && field[1][2] == "X" && field[2][2] == "X" ||
		field[0][0] == "X" && field[2][1] == "X" && field[2][2] == "X" {
		return 1
	} else if field[0][0] == "O" && field[0][1] == "O" && field[0][2] == "O" ||
		field[1][0] == "O" && field[1][1] == "O" && field[1][2] == "O" ||
		field[2][0] == "O" && field[2][1] == "O" && field[2][2] == "O" ||
		field[0][0] == "O" && field[1][0] == "O" && field[2][0] == "O" ||
		field[0][1] == "O" && field[1][1] == "O" && field[2][1] == "O" ||
		field[0][2] == "O" && field[1][2] == "O" && field[2][2] == "O" ||
		field[0][0] == "O" && field[2][1] == "O" && field[2][2] == "O" {
		return 0
	}
	return -1
}

func makeTurn(turn *int, field *[3][3]string) {
	var x, y int
	if *turn == 0 {
		*turn = 1
		fmt.Println("Player 1, your turn")
		fmt.Println("Enter cords to place 'O'")
		fmt.Print("Enter X cord: ")
		fmt.Scan(&x)
		fmt.Print("Enter Y cord: ")
		fmt.Scan(&y)
		if field[x][y] != "_" {
			fmt.Println("You cant override other symbol!")
			*turn = 0
			return
		}
		field[x][y] = "O"
		return
	} else if *turn == 1 {
		*turn = 0
		fmt.Println("Player 2, your turn")
		fmt.Println("Enter cords to place 'X'")
		fmt.Print("Enter X cord: ")
		fmt.Scan(&x)
		fmt.Print("Enter Y cord: ")
		fmt.Scan(&y)
		if field[x][y] != "_" {
			fmt.Println("You cant override other symbol!")
			*turn = 1
			return
		}
		field[x][y] = "X"
		return
	} else {
		panic("Invalid turn value")
	}
}

func displayField(field [3][3]string) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			fmt.Printf("%s ", field[i][j])
		}
		fmt.Println()
	}
}

func createField() [3][3]string {
	return [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}
