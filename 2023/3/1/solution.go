package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrixSquare struct {
	number int
	symbol bool
}

func (m *matrixSquare) isEmpty() bool {
	return m.number == -1 && !m.symbol
}

func (m *matrixSquare) isNumber() bool {
	return m.number != -1
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(data), "\n")

	num_cols := len(split[0])
	num_rows := len(split)
	matrix := make([][]matrixSquare, num_rows)
	for i := 0; i < num_rows; i++ {
		matrix[i] = make([]matrixSquare, num_cols)
	}

	buildMatrix(num_rows, split, matrix)
	total := solve(num_rows, num_cols, matrix)

	fmt.Println(total)
}

func solve(num_rows int, num_cols int, matrix [][]matrixSquare) int {
	currentNumber := 0
	total := 0
	for i := 0; i < num_rows; i++ {
		number_start_col := -1
		number_end_col := -1
		currentNumber = 0
		for j := 0; j < num_cols; j++ {
			square := matrix[i][j]

			if square.isNumber() {
				if number_start_col == -1 {
					number_start_col = j
				}
				currentNumber *= 10
				currentNumber += square.number
			}

			if number_start_col != -1 && (square.isEmpty() || square.symbol || j == num_cols-1) {
				// Handle number ending at the last column
				if j == num_cols-1 && square.isNumber() {
					number_end_col = j
				} else {
					number_end_col = j - 1
				}
				if isValidNumber(i, number_start_col, number_end_col, num_rows, num_cols, matrix) {
					// fmt.Println("Adding ", currentNumber)
					total += currentNumber
				}
				currentNumber = 0
				number_start_col = -1
				number_end_col = -1
			}
		}
	}
	return total
}

func isValidNumber(row int, start_col int, end_col int, num_rows int, num_cols int, matrix [][]matrixSquare) bool {
	cur_row := row - 1
	if cur_row >= 0 {
		// Check row above number
		rowHasSymbol := doesRowHaveSymbol(start_col, end_col, num_cols, matrix, cur_row)
		if rowHasSymbol {
			return true
		}
	}
	cur_row = row + 1
	if cur_row <= num_rows-1 {
		// Check row below number
		rowHasSymbol := doesRowHaveSymbol(start_col, end_col, num_cols, matrix, cur_row)
		if rowHasSymbol {
			return true
		}
	}

	cur_row = row
	// Check left end
	cur_col := start_col - 1
	if cur_col >= 0 {
		if matrix[cur_row][cur_col].symbol {
			return true
		}
	}

	// Check right end
	cur_col = end_col + 1
	if cur_col <= num_cols-1 {
		if matrix[cur_row][cur_col].symbol {
			return true
		}
	}

	return false
}

func doesRowHaveSymbol(start_col int, end_col int, num_cols int, matrix [][]matrixSquare, cur_row int) bool {
	for i := start_col - 1; i <= end_col+1; i++ {
		if i >= 0 && i <= num_cols-1 {
			if matrix[cur_row][i].symbol {
				return true
			}
		}
	}
	return false
}

func buildMatrix(num_rows int, split []string, matrix [][]matrixSquare) {
	for i := 0; i < num_rows; i++ {
		row := split[i]
		if row == "" {
			continue
		}
		for j := range row {
			char := string(row[j])
			square := &matrix[i][j]

			num, err := strconv.Atoi(char)
			if err == nil {
				square.number = num
				square.symbol = false
			} else if char == "." {
				square.number = -1
				square.symbol = false
			} else {
				square.number = -1
				square.symbol = true
			}
		}
	}
}
