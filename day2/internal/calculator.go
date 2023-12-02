package internal

import (
	"strconv"
	"strings"
)

func SumGamePowers(input []string) int {
	total := 0
	for _, line := range input {
		redMax := 0
		greenMax := 0
		blueMax := 0

		game := strings.Split(line, ":")[1]
		rounds := strings.Split(game, ";")
		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				trimmed := strings.Trim(cube, " ")
				tokens := strings.Split(trimmed, " ")
				amount, _ := strconv.Atoi(tokens[0])
				color := tokens[1]
				switch color {
				case "red":
					if amount > redMax {
						redMax = amount
					}
				case "green":
					if amount > greenMax {
						greenMax = amount
					}
				case "blue":
					if amount > blueMax {
						blueMax = amount
					}
				}
			}
		}

		power := redMax * greenMax * blueMax
		total += power
	}

	return total
}
