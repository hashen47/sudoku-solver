package sudoku_solver

import (
	"github.com/hashen47/sudoku-solver/board_lookup"
	"github.com/hashen47/sudoku-solver/dimension_lookup"
)

const (
	ROWS    = 9
	COLS    = 9
	MIN_VAL = 1
	MAX_VAL = 9
)

func GetSudokuSolutions(board Board) ([]Board, error) {
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
					return solutions, err
				}
			}
		}
	}

	err := solve(&board, &boardLookUpTable, &dimLookUpTable, &solutions)
	if err != nil {
		return solutions, err
	}

	return solutions, nil
}

func solve(board *Board, boardLookUpTable *board_lookup.BoardLookUp, dimLookUpTable *dimension_lookup.DimLookUp, solutions *[]Board) error {
	for rowI := 0; rowI < ROWS; rowI++ {
		for colI := 0; colI < COLS; colI++ {
			if (*board)[rowI][colI] == 0 {
				for val := MIN_VAL; val <= MAX_VAL; val++ {

					possible, err := isPossible(val, rowI, colI, boardLookUpTable, dimLookUpTable)
					if err != nil {
						return err
					}

					if possible {
						(*board)[rowI][colI] = val
						if err := boardLookUpTable.Insert(rowI, colI, val, true); err != nil {
							return err
						}
						dimLookUpTable.Insert(dimension_lookup.DimRow, rowI, val, true)
						dimLookUpTable.Insert(dimension_lookup.DimCol, colI, val, true)

						err := solve(board, boardLookUpTable, dimLookUpTable, solutions)
						if err != nil {
							return err
						}

						(*board)[rowI][colI] = 0
						if err := boardLookUpTable.Remove(rowI, colI, val); err != nil {
							return err
						}
						dimLookUpTable.Remove(dimension_lookup.DimRow, rowI, val)
						dimLookUpTable.Remove(dimension_lookup.DimCol, colI, val)
					}
				}
				return nil
			}
		}
	}

	*solutions = append(*solutions, board.Clone())

	return nil
}

func isPossible(val, row, col int, bl *board_lookup.BoardLookUp, dl *dimension_lookup.DimLookUp) (bool, error) {
	isExists, err := bl.IsKeyExists(row, col, val)
	if err != nil {
		return false, err
	}

	if isExists {
		return false, nil
	}

	if dl.IsKeyExists(dimension_lookup.DimRow, row, val) {
		return false, nil
	}

	if dl.IsKeyExists(dimension_lookup.DimCol, col, val) {
		return false, nil
	}

	return true, nil
}
