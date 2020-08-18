package sudoku

import "fmt"

// 0 represents an empty field
type Sudoku struct {
	Fields [9][9]int
}

func (sudoku *Sudoku) Print() {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			fmt.Printf(" %d ", sudoku.Fields[y][x])
		}
		fmt.Printf("\n")
	}
}

func (sudoku *Sudoku) Solve() error {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if sudoku.Fields[x][y] != 0 {
				continue
			}
			for n := 1; n < 10; n++ {
				if sudoku.possible(x, y, n) {
					sudoku.Fields[x][y] = n
					err := sudoku.Solve()
					if err != nil {
						sudoku.Fields[x][y] = 0
					} else {
						return nil
					}
				}
			}
			return fmt.Errorf("dead end")
		}
	}

	return nil
}

func (sudoku *Sudoku) possible(x, y, n int) bool {
	for gx := 0; gx < 9; gx++ {
		if sudoku.Fields[gx][y] == n {
			return false
		}
	}
	for gy := 0; gy < 9; gy++ {
		if sudoku.Fields[x][gy] == n {
			return false
		}
	}

	bx := (x / 3) * 3
	by := (y / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku.Fields[bx+i][by+j] == n {
				return false
			}
		}
	}

	return true
}
