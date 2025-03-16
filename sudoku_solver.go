package sudoku_solver

import (
	"fmt"
	"github.com/hashen47/sudoku-solver/dimension_lookup"
	"github.com/hashen47/sudoku-solver/board_lookup"
)

const (
	ROWS = 9 
	COLS = 9 
	MIN_VAL = 1 
	MAX_VAL = 9 
)

func main() {
	board := Board{
		[]int{5, 3, 0, 0, 7, 0, 0, 0, 0},
		[]int{6, 0, 0, 1, 9, 5, 0, 0, 0},
		[]int{0, 9, 8, 0, 0, 0, 0, 6, 0},
		[]int{8, 0, 0, 0, 6, 0, 0, 0, 3},
		[]int{4, 0, 0, 8, 0, 3, 0, 0, 1},
		[]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
		[]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
		[]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
		[]int{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solutions := GetSudokuSolutions(board)

	for _, solution := range solutions {
		fmt.Println()
		fmt.Println(solution)
		fmt.Println()
	}
}

func GetSudokuSolutions(board Board) []Board {
	solutions := make([]Board, 0)
	boardLookUpTable := board_lookup.NewBoardLookUp()
	dimLookUpTable := dimension_lookup.NewDimLookUp()

	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			value := (board)[row][col]
			if value != 0 {
				dimLookUpTable.Insert(dimension_lookup.DimRow, row, value, true)
				dimLookUpTable.Insert(dimension_lookup.DimCol, col, value, true)
				if err := boardLookUpTable.Insert(row, col, value, true); err != nil {
					panic(err)
				}
			}
		}
	}

	solve(&board, &boardLookUpTable, &dimLookUpTable, &solutions)

	return solutions
}

func solve(board *Board, boardLookUpTable *board_lookup.BoardLookUp, dimLookUpTable *dimension_lookup.DimLookUp, solutions *[]Board) {
	for rowI := 0; rowI < ROWS; rowI++ {
		for colI := 0; colI < COLS; colI++ {
			if (*board)[rowI][colI] == 0 {
				for val := MIN_VAL; val <= MAX_VAL; val++ {
					if isPossible(val, rowI, colI, boardLookUpTable, dimLookUpTable) {
						(*board)[rowI][colI] = val
						if err := boardLookUpTable.Insert(rowI, colI, val, true); err != nil {
							panic(err)
						}
						dimLookUpTable.Insert(dimension_lookup.DimRow, rowI, val, true)
						dimLookUpTable.Insert(dimension_lookup.DimCol, colI, val, true)

						solve(board, boardLookUpTable, dimLookUpTable, solutions)

						(*board)[rowI][colI] = 0
						if err := boardLookUpTable.Remove(rowI, colI, val); err != nil {
							panic(err)
						}
						dimLookUpTable.Remove(dimension_lookup.DimRow, rowI, val)
						dimLookUpTable.Remove(dimension_lookup.DimCol, colI, val)
					}
				}
				return
			}
		}
	}

	*solutions = append(*solutions, copyBoard(board))
}

func isPossible(val, row, col int, bl *board_lookup.BoardLookUp, dl *dimension_lookup.DimLookUp) bool {
	isExists, err := bl.IsKeyExists(row, col, val)
	if err != nil {
		panic(err)
	}

	if isExists {
		return false
	}

	if dl.IsKeyExists(dimension_lookup.DimRow, row, val) {
		return false
	}

	if dl.IsKeyExists(dimension_lookup.DimCol, col, val) {
		return false
	}

	return true
}

func copyBoard(srcBoard *Board) Board {
	board := make(Board, 0)
	for i := 0; i < len(*srcBoard); i++ {
		row := make([]int, 0)
		for j := 0; j < len((*srcBoard)[0]); j++ {
			row = append(row, (*srcBoard)[i][j]) 
		}
		board = append(board, row)
	}
	return board
}
