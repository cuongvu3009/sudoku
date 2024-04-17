# Sudoku Solver in Go

This Go program solves Sudoku puzzles using a backtracking algorithm. It takes a 9x9 Sudoku grid as input and prints the solved grid.

## Overview

The program takes a Sudoku grid as input, either as command-line arguments or as a single string without spaces. It validates the input grid and then solves the Sudoku puzzle recursively using a backtracking algorithm. Finally, it prints the solved Sudoku grid.

## Components

1. **Main Function (`main`)**:
   - Parses command-line arguments or a single string input.
   - Validates the input grid.
   - Solves the Sudoku puzzle and prints the solution.

2. **Sudoku Grid Type**:
   - Represents a 9x9 Sudoku grid as a 2D array of integers.

3. **isValid Function**:
   - Checks if the given Sudoku grid is valid according to Sudoku rules.
   - Ensures that each row, column, and 3x3 subgrid contains unique integers from 1 to 9.

4. **solveSudoku Function**:
   - Recursively solves the Sudoku puzzle using a backtracking algorithm.
   - Attempts to fill in empty cells with valid numbers and backtracks if a solution cannot be found.

5. **isValidPlacement Function**:
   - Checks if placing a number at a specific cell in the Sudoku grid is valid.
   - Ensures that the number does not violate Sudoku rules in its row, column, or 3x3 subgrid.

6. **SplitWhiteSpaces Function**:
   - Splits a string into substrings based on white spaces (spaces, tabs, newlines).

7. **printGrid Function**:
   - Prints the Sudoku grid to the console.

## How to Run

1. Ensure you have Go installed on your system.
2. Save the provided code in a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the following command to execute the program:

   ```bash
   go run . "..213.748" "8.4.....2" ".178.26.." ".68.9.27." ".932....4" "5..46.3.." "..9.24.23" "..63..19." "385..1.2."
   ```

## Alternatively, you can build the program and execute the binary:

```bash
go build .
./main "..213.748" "8.4.....2" ".178.26.." ".68.9.27." ".932....4" "5..46.3.." "..9.24.23" "..63..19." "385..1.2."
```
