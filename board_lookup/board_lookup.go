package board_lookup

import (
	"errors"
)

type BoardLookUpType int
const (
	BoardOne BoardLookUpType = iota 
	BoardTwo
	BoardThre
	BoardFour
	BoardFive
	BoardSix
	BoardSeven
	BoardEight
	BoardNine
)

type BoardLookUp map[BoardLookUpType]map[int]bool 

var LookUpOutBoundError error = errors.New("row or column value cannot less than 0 or cannot exceed 8")
var LookUpKeyNotFoundError error = errors.New("given key is not found in the lookup")

func NewBoardLookUp() BoardLookUp {
	l := make(BoardLookUp, 0)
	for i := 0; i < 9; i++ {
		l[BoardLookUpType(i)]   = make(map[int]bool, 0)
	}
	return l
}

func (l *BoardLookUp) Insert(row, col , key int, val bool) error {
	if !isRowColRangeValid(row, col) {
		return LookUpOutBoundError
	}
	boardType := getBoardLookUpType(row, col)
	(*l)[boardType][key] = val
	return nil
}

func (l *BoardLookUp) IsKeyExists(row, col, key int) (bool, error) {
	if !isRowColRangeValid(row, col) {
		return false, LookUpOutBoundError
	}
	boardType := getBoardLookUpType(row, col)
	_, ok := (*l)[boardType][key]
	return ok, nil
}

func (l *BoardLookUp) Remove(row, col, key int) error {
	isExists, err := l.IsKeyExists(row, col, key)
	if err != nil {
		return err
	}

	if !isExists {
		return LookUpKeyNotFoundError
	}

	boardType := getBoardLookUpType(row, col)
	delete((*l)[boardType], key)

	return nil
}

func isRowColRangeValid(row, col int) bool {
	return row >= 0 && col >= 0 && row <= 8 && col <= 8
}

func getBoardLookUpType(row, col int) BoardLookUpType {
	return BoardLookUpType((row/3)*3 + (col/3))
}
