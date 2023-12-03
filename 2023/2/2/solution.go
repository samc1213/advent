package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPower(row string) (int, error) {
	necessary_per_color := map[string]int{"red": 0, "green": 0, "blue": 0}
	rowSplit := strings.Split(row, ":")
	if len(rowSplit) != 2 {
		return 0, fmt.Errorf("Expected only one colon. Given %s", row)
	}

	drawsSplit := strings.Split(rowSplit[1], ";")
	for i := range drawsSplit {
		draw := drawsSplit[i]
		drawSplit := strings.Split(draw, ",")
		for j := range drawSplit {
			numberColor := drawSplit[j]
			numberColor = strings.Trim(numberColor, " ")
			numberColorSplit := strings.Split(numberColor, " ")
			number, err := strconv.Atoi(numberColorSplit[0])
			if err != nil {
				return 0, err
			}
			color := numberColorSplit[1]
			if necessary_per_color[color] < number {
				necessary_per_color[color] = number
			}
		}
	}

	power := 1

	for color := range necessary_per_color {
		power *= necessary_per_color[color]
	}

	return power, nil
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
		power, err := getPower(row)
		if err != nil {
			panic(err)
		}
		total += power
	}

	fmt.Println(total)
}
