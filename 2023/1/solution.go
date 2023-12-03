package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getDigit(row string, index int) (int, error) {
	digit_str_to_num := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	for digit_str := range digit_str_to_num {
		if strings.HasPrefix(row[index:], digit_str) {
			return digit_str_to_num[digit_str], nil
		}
	}

	return 0, fmt.Errorf("No digit found")
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(data), "\n")

	var total int
	for i := range split {
		row := split[i]
		if row == "" {
			continue
		}
		var first_num int
		var last_num int
		for j := 0; j < len(row); j++ {
			num, err := strconv.Atoi(string(row[j]))

			if err != nil {
				num, err = getDigit(row, j)
				if err != nil {
					continue
				}
			}

			first_num = num
			break
		}

		for j := len(row) - 1; j >= 0; j-- {
			num, err := strconv.Atoi(string(row[j]))

			if err != nil {
				num, err = getDigit(row, j)
				if err != nil {
					continue
				}
			}
			last_num = num
			break
		}

		row_total, err := strconv.Atoi(strconv.Itoa(first_num) + strconv.Itoa(last_num))
		if err != nil {
			panic(err)
		}
		total += row_total
	}

	fmt.Println(total)
}
