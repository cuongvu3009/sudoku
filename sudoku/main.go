package main

import (
	"fmt"
	"os"
)

type SudokuGrid [9][9]int

func main() {
	args := os.Args[1:]

	var str string
	if len(args) == 1 && len(args[0]) != 0 {
		for _, c := range args[0] {
			if c != '"' {
				str = str + string(c)
			}
		}
		args = SplitWhiteSpaces(str)
	}

	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	var grid SudokuGrid
	for c, row := range args {
		if len(row) != 9 {
			fmt.Printf("Error")
			return
		}
		for r := 0; r < 9; r++ {
			ch := row[r]
			if ch == '.' {
				grid[c][r] = 0
			} else if ch >= '1' && ch <= '9' {
				grid[c][r] = int(ch - '0')
			} else {
				fmt.Printf("Error")
				return
			}
		}
	}

	if !isValid(grid) {
		fmt.Println("Error")
		return
	}

	if !solveSudoku(&grid) {
		fmt.Println("Error")
		return
	}

	printGrid(grid)
}

func isValid(grid SudokuGrid) bool {
	for c := 0; c < 9; c++ {
		nums := make(map[int]bool)
		for r := 0; r < 9; r++ {
			num := grid[c][r]
			if num != 0 && nums[num] {
				return false
			} else if num != 0 {
				nums[num] = true
			}
		}
	}
	for r := 0; r < 9; r++ {
		nums := make(map[int]bool)
		for c := 0; c < 9; c++ {
			num := grid[c][r]
			if num != 0 && nums[num] {
				return false
			} else if num != 0 {
				nums[num] = true
			}
		}
	}
	for c := 0; c < 9; c += 3 {
		for r := 0; r < 9; r += 3 {
			nums := make(map[int]bool)
			for x := c; x < c+3; x++ {
				for y := r; y < r+3; y++ {
					num := grid[x][y]
					if num != 0 && nums[num] {
						return false
					} else if num != 0 {
						nums[num] = true
					}
				}
			}
		}
	}
	return true
}

func solveSudoku(grid *SudokuGrid) bool {
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			if grid[c][r] == 0 {
				for num := 1; num <= 9; num++ {
					if isValidPlacement(grid, c, r, num) {
						grid[c][r] = num
						if solveSudoku(grid) {
							return true
						}
						grid[c][r] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isValidPlacement(grid *SudokuGrid, c, r, num int) bool {
	secI := c / 3 * 3
	secJ := r / 3 * 3
	for x := secI; x < secI+3; x++ {
		for y := secJ; y < secJ+3; y++ {
			if grid[x][y] == num {
				return false
			}
		}
	}
	for x := 0; x < 9; x++ {
		if grid[c][x] == num || grid[x][r] == num {
			return false
		}
	}
	return true
}
func SplitWhiteSpaces(s string) []string {
	var result []string

	var wordStart int

	for c := 0; c < len(s); c++ {
		if s[c] == ' ' || s[c] == '\t' || s[c] == '\n' {
			if wordStart < c {
				result = append(result, s[wordStart:c])
			}
			wordStart = c + 1
		}
	}
	if wordStart < len(s) {
		result = append(result, s[wordStart:])
	}
	return result
}

func printGrid(grid SudokuGrid) {
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			fmt.Print(" ", grid[c][r])
		}
		fmt.Println()
	}
}
