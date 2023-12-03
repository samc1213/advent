package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrixSquare struct {
	number               int
	symbol               string
	adjacent_partNumbers []int
}

func (m *matrixSquare) isEmpty() bool {
	return m.number == -1 && m.symbol == ""
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

			if number_start_col != -1 && (square.isEmpty() || square.symbol != "" || j == num_cols-1) {
				// Handle number ending at the last column
				if j == num_cols-1 && square.isNumber() {
					number_end_col = j
				} else {
					number_end_col = j - 1
				}
				addAdjacentPartNumberToStars(currentNumber, i, number_start_col, number_end_col, num_rows, num_cols, matrix)
				currentNumber = 0
				number_start_col = -1
				number_end_col = -1
			}
		}
	}

	for i := 0; i < num_rows; i++ {
		for j := 0; j < num_cols; j++ {

			if len(matrix[i][j].adjacent_partNumbers) == 2 {
				total += matrix[i][j].adjacent_partNumbers[0] * matrix[i][j].adjacent_partNumbers[1]
			}
		}
	}

	return total
}

func addAdjacentPartNumberToStars(currentNumber int, row int, start_col int, end_col int, num_rows int, num_cols int, matrix [][]matrixSquare) {
	cur_row := row - 1
	if cur_row >= 0 {
		// Check row above number
		addAdjacentPartNumberForStarsInRow(currentNumber, start_col, end_col, num_cols, matrix, cur_row)
	}
	cur_row = row + 1
	if cur_row <= num_rows-1 {
		// Check row below number
		addAdjacentPartNumberForStarsInRow(currentNumber, start_col, end_col, num_cols, matrix, cur_row)
	}

	cur_row = row
	// Check left end
	cur_col := start_col - 1
	if cur_col >= 0 {
		if matrix[cur_row][cur_col].symbol == "*" {
			matrix[cur_row][cur_col].adjacent_partNumbers = append(matrix[cur_row][cur_col].adjacent_partNumbers, currentNumber)
		}
	}

	// Check right end
	cur_col = end_col + 1
	if cur_col <= num_cols-1 {
		if matrix[cur_row][cur_col].symbol == "*" {
			matrix[cur_row][cur_col].adjacent_partNumbers = append(matrix[cur_row][cur_col].adjacent_partNumbers, currentNumber)
		}
	}
}

func addAdjacentPartNumberForStarsInRow(currentNumber int, start_col int, end_col int, num_cols int, matrix [][]matrixSquare, cur_row int) {
	for i := start_col - 1; i <= end_col+1; i++ {
		if i >= 0 && i <= num_cols-1 {
			if matrix[cur_row][i].symbol == "*" {
				matrix[cur_row][i].adjacent_partNumbers = append(matrix[cur_row][i].adjacent_partNumbers, currentNumber)
			}
		}
	}
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
				square.symbol = ""
			} else if char == "." {
				square.number = -1
				square.symbol = ""
			} else {
				square.number = -1
				square.symbol = char
			}
		}
	}
}
