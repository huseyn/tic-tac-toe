package main

import (
	"fmt"
	"strings"
)

func main() {
	var row, col int
	var moveCount int = 0
	var val string
	const totalMoves int = 9

	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	for {
		displayBoard(board)
		fmt.Print("Enter row, col and value: ")

		if _, err := fmt.Scanln(&row, &col, &val); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Row and column must be between 0 and 2.")
			continue
		}
		if val != "X" && val != "O" {
			fmt.Println("Value must be either 'X' or 'O'.")
			continue
		}
		if board[row][col] != "_" {
			fmt.Println("Position already taken, please choose another.")
			continue
		}

		board[row][col] = val
		moveCount++

		winner := checkWinner(board)
		if winner != "" {
			fmt.Println("Winner:", winner)
			displayBoard(board)
			break
		} else if moveCount == totalMoves {
			fmt.Println("Draw!")
			displayBoard(board)
			break
		}
	}
}

func displayBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func checkWinner(board [][]string) string {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] != "_" && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != "_" && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return board[0][i]
		}
	}

	// Check diagonals
	if board[0][0] != "_" && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != "_" && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}

	// No winner found
	return ""
}
