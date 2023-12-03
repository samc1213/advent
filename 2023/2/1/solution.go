package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getIdtoAdd(row string) (int, error) {
	max_per_color := map[string]int{"red": 12, "green": 13, "blue": 14}
	rowSplit := strings.Split(row, ":")
	if len(rowSplit) != 2 {
		return 0, fmt.Errorf("Expected only one colon. Given %s", row)
	}
	r, err := regexp.Compile("Game ([0-9]*)")
	if err != nil {
		return 0, errors.New("Bad regex")
	}
	game, err := strconv.Atoi(r.FindStringSubmatch(rowSplit[0])[1])
	if err != nil {
		return 0, nil
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
			if max_per_color[color] < number {
				return 0, nil
			}
		}
	}
	return game, nil
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
		id, err := getIdtoAdd(row)
		if err != nil {
			panic(err)
		}
		total += id
	}

	fmt.Println(total)
}
