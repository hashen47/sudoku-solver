package dimension_lookup

import (
	"testing"
)

func TestInsert(t *testing.T) {
	rows, cols := 9, 9
	dl := NewDimLookUp()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dl.Insert(DimRow, i, j, true)
			dl.Insert(DimCol, i, j, true)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := dl[DimRow][i][j]; !ok {
				t.Fatalf("ERROR (DimRow): dl: %v, expect: true,", dl)
			}
			if _, ok := dl[DimCol][i][j]; !ok {
				t.Fatalf("ERROR (DimCol): dl: %v, expect: true,", dl)
			}
		}
	}
}

func TestIsKeyExists(t *testing.T) {
	rows, cols := 9, 9
	dl := NewDimLookUp()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dl.Insert(DimRow, i, j, true)
			dl.Insert(DimCol, i, j, true)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !dl.IsKeyExists(DimRow, i, j) {
				t.Fatalf("ERROR (DimRow): dl: %v, expect: true,", dl)
			}
			if !dl.IsKeyExists(DimCol, i, j) {
				t.Fatalf("ERROR (DimCol): dl: %v, expect: true,", dl)
			}
		}
	}

	for i := rows; i < rows*2; i++ {
		for j := 0; j < cols; j++ {
			if dl.IsKeyExists(DimRow, i, j) {
				t.Fatalf("ERROR (DimRow): dl: %v, expect: true,", dl)
			}
			if dl.IsKeyExists(DimCol, i, j) {
				t.Fatalf("ERROR (DimCol): dl: %v, expect: true,", dl)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := cols; j < cols*2; j++ {
			if dl.IsKeyExists(DimRow, i, j) {
				t.Fatalf("ERROR (DimRow): dl: %v, expect: true,", dl)
			}
			if dl.IsKeyExists(DimCol, i, j) {
				t.Fatalf("ERROR (DimCol): dl: %v, expect: true,", dl)
			}
		}
	}
}
