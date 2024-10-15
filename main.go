package main

import (
	"fmt"

	p "github.com/jlfjunior/learning-golang/problems"
)

func main() {
	board := make([][]int, 2)

	for i := range board {
		board[i] = make([]int, len(board))
	}

	fmt.Println("Imprimir tabuleiro")
	p.PrintBoard(board)

	if p.SolverQueensPuzzle(board, 0, len(board)) {

		fmt.Println("Imprimir resultado")
		p.PrintBoard(board)
	} else {
		fmt.Println("Solução não encontrada!")
	}
}
