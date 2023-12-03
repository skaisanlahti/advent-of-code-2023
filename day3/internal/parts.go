package internal

import (
	"math"
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
	parts := parseEnginePartsWithMath(&schema)
	return calculateNumberTotal(parts, &schema)
}

func parseEnginePartsWithRegex(input []string) []enginePart {
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

func parseEnginePartsWithMath(s *schema) []enginePart {
	var parts []enginePart
	for row, runes := range s.content {
		for column := 0; column < s.columns; {
			digit, ok := runeToInt(runes[column])
			if !ok {
				column++
				continue
			}

			digits := []int{digit}
			for next := column + 1; next < s.columns; next++ {
				digit, ok := runeToInt(runes[next])
				if !ok {
					break
				}

				digits = append(digits, digit)
			}

			length := len(digits)
			number := 0
			for i, m := 0, length-1; i < length; i, m = i+1, m-1 {
				number += digits[i] * int(math.Pow10(m))
			}

			part := enginePart{number, row, column, length}
			parts = append(parts, part)

			column += length
		}
	}

	return parts
}

func runeToInt(r rune) (int, bool) {
	if r < '0' || r > '9' {
		return 0, false
	}

	return int(r - '0'), true
}

func calculateNumberTotal(parts []enginePart, schema *schema) int {
	total := 0
	for _, part := range parts {
		neighbors := findNeighbors(part, schema)
		for _, neighbor := range neighbors {
			if isSymbol(neighbor.symbol) {
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
