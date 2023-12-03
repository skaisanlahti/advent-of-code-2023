package internal

import (
	"regexp"
	"strconv"
)

type enginePart struct {
	number int
	row    int
	column int
	length int
}

func SumPartNumbers(input []string) int {
	schema := newSchema(input)
	parts := findEngineParts(input)
	return calculateNumberTotal(parts, &schema)
}

func findEngineParts(input []string) []enginePart {
	var parts []enginePart
	for row, line := range input {
		finder := regexp.MustCompile(`[0-9]+`)
		partStrings := finder.FindAllString(line, -1)
		partStringIndexes := finder.FindAllStringIndex(line, -1)

		for i, partString := range partStrings {
			number, _ := strconv.Atoi(partString)
			column := partStringIndexes[i][0]
			length := len([]rune(partString))
			part := enginePart{number, row, column, length}
			parts = append(parts, part)
		}
	}

	return parts
}

func calculateNumberTotal(parts []enginePart, schema *schema) int {
	total := 0
	for _, part := range parts {
		neighbors := findNeighbors(part, schema)
		for _, neighbor := range neighbors {
			ok := isSymbol(neighbor.char)
			if ok {
				total += part.number
				break
			}
		}

	}

	return total
}

func isSymbol(r rune) bool {
	if r == '.' {
		return false
	}

	if r >= '0' && r <= '9' {
		return false
	}

	return true
}
