package sudoku_solver

import (
	"fmt"
)

type Board [][]int

func NewBoard(rows, cols int) Board {
	board := make(Board, 0)
	for rowI := 0; rowI < rows; rowI++ {
		row := make([]int, cols)
		board = append(board, row)
	}
	return board
}

func (b *Board) Clone() Board {
	boardCopy := make(Board, 0)
	for row := 0; row < len(*b); row++ {
		rowCopy := make([]int, 0)
		rowCopy = append(rowCopy, (*b)[row]...)
		boardCopy = append(boardCopy, rowCopy)
	}
	return boardCopy
}

func (b *Board) IsEqual(b2 *Board) bool {
	if len(*b) != len(*b2) {
		return false
	}

	for row := 0; row < len(*b); row++ {
		if len((*b)[row]) != len((*b2)[row]) {
			return false
		}

		for col := 0; col < len((*b)[row]); col++ {
			if (*b)[row][col] != (*b2)[row][col] {
				return false
			}
		}
	}

	return true
}

func (b Board) String() string {
	output := ""
	for rowI, row := range b {
		for _, col := range row {
			output += fmt.Sprintf("%d ", col)
		}
		if rowI != len(b)-1 {
			output += "\n"
		}
	}
	return output
}
