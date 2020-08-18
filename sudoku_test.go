package sudoku_test

import (
	"testing"

	sudoku "github.com/pako8128/sudoku_solver"
)

func TestSudoku1(t *testing.T) {
	s := sudoku.Sudoku {
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

	if err := s.Solve(); err != nil {
		s.Print()
		t.Fatalf("Could not Solve Suduko, got error: %v", err)
	}
}

func TestSudoku2(t *testing.T) {
	s := sudoku.Sudoku {
		Fields: [9][9]int{
			{ 1, 2, 3, 0, 0, 0, 0, 0, 0 },
			{ 4, 5, 6, 0, 0, 0, 0, 0, 0 },
			{ 7, 8, 9, 0, 0, 0, 0, 0, 0 },
			{ 0, 0, 0, 0, 0, 0, 0, 0, 0 },
			{ 0, 0, 0, 0, 0, 0, 0, 0, 0 },
			{ 0, 0, 0, 0, 0, 0, 0, 0, 0 },
			{ 0, 0, 0, 0, 0, 0, 1, 2, 3 },
			{ 0, 0, 0, 0, 0, 0, 4, 5, 6 },
			{ 0, 0, 0, 0, 0, 0, 7, 8, 9 },
		},
	}

	if err := s.Solve(); err != nil {
		s.Print()
		t.Fatalf("Could not Solve Suduko, got error: %v", err)
	}
}

func TestUnsolvableSudoku(t *testing.T) {
	s := sudoku.Sudoku {
		Fields: [9][9]int{
			{ 3, 4, 6, 0, 0, 9, 0, 0, 1 },
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

	if err := s.Solve(); err == nil {
		t.Fatal("Solved the unsolvable Sudoku")
	}
}

func TestSolveSudokuTwice(t *testing.T) {
	s := sudoku.Sudoku {
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

	_ = s.Solve()
	if err := s.Solve(); err != nil {
		t.Fatalf("Could not Solve already solved Sudoku, got error: %v", err)
	}
}


