package main

import (
    "fmt"
    "strings"
)

func createBoard() [][]string {
    board := [][]string{
        {"_", "_", "_"},
        {"_", "_", "_"},
        {"_", "_", "_"},
    }

    return board
}

func printBoard(board [][]string) {
	fmt.Println("Board:")
	for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}

func playerInput(board [][]string, symbol string) {
    var input int
    for {
        fmt.Print("Enter your move (1-9): ")
        _, err := fmt.Scan(&input)
        if err != nil {
            fmt.Println("Error reading input. Please enter a number between 1 and 9.")
            fmt.Scanln()
            continue
        }
        if input < 1 || input > 9 {
            fmt.Println("Invalid input. Please enter a number between 1 and 9.")
            continue
        }

        row, col := (input-1)/3, (input-1)%3

        if board[row][col] != "_" {
            fmt.Println("This position is already taken. Please try another one.")
            continue
        }

        board[row][col] = symbol
        break
    }
}

func checkWin(board [][]string, symbol string) bool {
	
	for i := 0; i < 3; i++ {
		if (board[i][0] == symbol && board[i][1] == symbol && board[i][2] == symbol) ||
			(board[0][i] == symbol && board[1][i] == symbol && board[2][i] == symbol) {
			return true
		}
	}

	if (board[0][0] == symbol && board[1][1] == symbol && board[2][2] == symbol) || 
		(board[0][2] == symbol && board[1][1] == symbol && board[2][0] == symbol) {
		return true
	}

	return false
}

func main() {
	fmt.Println("Welcome to tic-tac-toe! Tic-tac-toe is a classic paper-and-pencil game where two players take turns marking spaces in a 3x3 grid with their respective symbols, typically X and O, aiming to form a row, column, or diagonal of their symbols before their opponent does.")
	board := createBoard()
	currentPlayer := "O"

	for {
		printBoard(board)
		playerInput(board, currentPlayer)
		if checkWin(board, currentPlayer) {
			fmt.Println(currentPlayer, "wins! ")
			break
		}

		if currentPlayer == "O" {
            currentPlayer = "X"
        } else {
            currentPlayer = "O"
        }
	}
	
}
