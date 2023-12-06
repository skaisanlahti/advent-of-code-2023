package cube

import (
	"strconv"
	"strings"
)

func SumValidGameIds(input []string, red, green, blue int) int {
	limits := map[string]int{
		"red":   red,
		"green": green,
		"blue":  blue,
	}

	total := 0
	for i, line := range input {
		game := strings.Split(line, ":")[1]
		rounds := strings.Split(game, ";")
		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				trimmed := strings.Trim(cube, " ")
				tokens := strings.Split(trimmed, " ")
				amount, _ := strconv.Atoi(tokens[0])
				color := tokens[1]
				limit, _ := limits[color]
				if amount > limit {
					// game is invalid, skip rest of processing
					goto EndGameValidation
				}
			}
		}

		// game is valid if reached here without jumping to EndGameValidation
		total += i + 1
	EndGameValidation:
	}

	return total
}
