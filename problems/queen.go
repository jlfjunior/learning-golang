package problems

import "fmt"

func SolverQueensPuzzle(board [][]int, colunm, n int) bool {

	if colunm >= n {
		return true
	}

	for row := 0; row < len(board); row++ {
		if isSafe(board, row, colunm) {
			board[row][colunm] = 1

			if SolverQueensPuzzle(board, colunm+1, n) {
				return true
			}

			//Here a backtracking
			board[row][colunm] = 0
		}
	}

	return false
}

func isSafe(board [][]int, row, column int) bool {
	hasQueen := 1
	//Check columns
	for c := 0; c < column; c++ {
		if board[row][c] == hasQueen {
			return false
		}
	}

	//Check upper Diagonal on left side

	for r, c := row, column; r >= 0 && c >= 0; {
		if board[r][c] == hasQueen {
			return false
		}

		r--
		c--
	}

	//Check down Diagonal on left side

	for r, c := row, column; r < len(board) && c >= 0; {
		if board[r][c] == hasQueen {
			return false
		}

		r++
		c--
	}

	return true
}

func PrintBoard(board [][]int) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
