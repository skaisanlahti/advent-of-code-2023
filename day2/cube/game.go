package cube

import (
	"strconv"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type Game struct {
	Id    int
	Cubes []Cube
}

type Cube struct {
	Color string
	Count int
}

func SumValidGameIds(input []string, red, green, blue int) int {
	limits := map[string]int{
		"red":   red,
		"green": green,
		"blue":  blue,
	}

	games := parseGames(input)
	sum := 0
	for _, game := range games {
		valid := true
		for _, cube := range game.Cubes {
			limit := limits[cube.Color]
			if cube.Count > limit {
				valid = false
				break
			}
		}

		if valid {
			sum += game.Id
		}
	}

	return sum
}

func SumGamePowers(input []string) int {
	games := parseGames(input)
	sum := 0
	for _, game := range games {
		redMax := 0
		greenMax := 0
		blueMax := 0

		for _, cube := range game.Cubes {
			switch cube.Color {
			case "red":
				if cube.Count > redMax {
					redMax = cube.Count
				}
			case "green":
				if cube.Count > greenMax {
					greenMax = cube.Count
				}
			case "blue":
				if cube.Count > blueMax {
					blueMax = cube.Count
				}
			}
		}

		power := redMax * greenMax * blueMax
		sum += power
	}

	return sum
}

func parseGames(input []string) []Game {
	games := []Game{}
	for _, line := range input {
		sections := strings.Split(line, ":")
		id := kit.NumbersFromString(sections[0])[0]
		rounds := strings.Split(sections[1], ";")
		cubes := []Cube{}
		for _, round := range rounds {
			cubesStrings := strings.Split(round, ",")
			for _, cubeString := range cubesStrings {
				cubeData := strings.Trim(cubeString, " ")
				tokens := strings.Split(cubeData, " ")
				amount, _ := strconv.Atoi(tokens[0])
				color := tokens[1]
				cubes = append(cubes, Cube{color, amount})
			}
		}

		games = append(games, Game{id, cubes})
	}

	return games
}
