package sudoku

import (
	"fmt"
	"testing"
)

func TestSudoku(t *testing.T) {
	sudoku := Sudoku {
		Fields: [9][9]int{
			{ 3, 4, 0, 0, 0, 9, 0, 0, 1 },
			{ 0, 0, 0, 2, 0, 1, 0, 0, 3 },
			{ 0, 6, 2, 0, 0, 0, 7, 8, 0 },
			{ 0, 0, 5, 0, 1, 8, 0, 0, 4 },
			{ 7, 9, 0, 4, 0, 3, 0, 5, 8 },
			{ 4, 0, 0, 7, 5, 0, 9, 0, 0 },
			{ 0, 7, 3, 0, 0, 0, 1, 9, 0 },
			{ 9, 0, 0, 1, 0, 2, 0, 0, 0 },
			{ 8, 0, 0, 6, 0, 0, 0, 3, 5 },
		},
	}

	sudoku.Print()
	_ = sudoku.Solve()

	fmt.Println("\n------------------------------")
	sudoku.Print()
}
