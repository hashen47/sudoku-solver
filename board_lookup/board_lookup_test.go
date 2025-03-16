package board_lookup

import (
	"testing"
)

func TestInsert(t *testing.T) {
	bl := NewBoardLookUp()

	row, col := 0, 0

	for i := 0; i < 100; i++ {
		if err := bl.Insert(row, col, i, true); err != nil {
			t.Fatalf("ERROR: row: %d, col: %d, inserted value: true, error(expect): nil, error(real): '%v'\n", row, col, err)
		}
	}

	for i := 0; i < 100; i++ {
		if bl[getBoardLookUpType(row, col)][i] != true {
			t.Fatalf("ERROR: row: %d, col: %d, expect: true, real: false\n", row, col)
		}
	}
}

func TestIsKeyExists(t *testing.T) {
	bl := NewBoardLookUp()

	row, col := 8, 8
	for i := 0; i < 100; i++ {
		if err := bl.Insert(row, col, i, true); err != nil {
			t.Fatalf("ERROR: row: %d, col: %d, inserted value: true, error(expect): nil, error(real): '%v'\n", row, col, err)
		}
	}

	for i := 0; i < 100; i++ {
		isExists, err := bl.IsKeyExists(row, col, i)
		if err != nil {
			t.Fatalf("ERROR: row: %d, col: %d, value: %v, error(expect): nil, error(real): '%v'\n", row, col, i, err)
		}
		if !isExists {
			t.Fatalf("ERROR: row: %d, col: %d, value: %v, expect: true, real: false\n", row, col, i)
		}
	}

	for i := 100; i < 200; i++ {
		isExists, err := bl.IsKeyExists(row, col, i)
		if err != nil {
			t.Fatalf("ERROR: row: %d, col: %d, value: %v, error(expect): nil, error(real): '%v'\n", row, col, i, err)
		}
		if isExists {
			t.Fatalf("ERROR: row: %d, col: %d, value: %v, expect: false, real: true\n", row, col, i)
		}
	}
}

func TestRemove(t *testing.T) {
	bl := NewBoardLookUp()

	row, col := 2, 7

	for i := 0; i < 100; i++ {
		if err := bl.Insert(row, col, i, true); err != nil {
			t.Fatalf("ERROR: row: %d, col: %d, inserted value: true, error(expect): nil, error(real): '%v'\n", row, col, err)
		}
	}

	for i := 0; i < 100; i++ {
		if err := bl.Remove(row, col, i); err != nil {
			t.Fatalf("ERROR: row: %d, col: %d, key: %v, error(expect): nil, error(real): '%v', lookup: %v\n", row, col, i, err, bl[getBoardLookUpType(row, col)][i])
		}
	}

	for i := 100; i < 200; i++ {
		if err := bl.Remove(row, col, i); err == nil {
			t.Fatalf("ERROR: row: %d, col: %d, error(expect): '%v', error(real): nil\n", row, col, LookUpKeyNotFoundError)
		}
	}
}
