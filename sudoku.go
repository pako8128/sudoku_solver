package sudoku

import "fmt"

// -1 represents an empty field
type Sudoku struct {
	Fields [9][9]int
}

func (sudoku *Sudoku) Solve() (error) {
	return fmt.Errorf("not implemented!")
}

