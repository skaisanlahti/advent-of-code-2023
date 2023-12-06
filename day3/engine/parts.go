package engine

import (
	"math"
	"regexp"
	"strconv"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type EnginePart struct {
	Number int
	Row    int
	Column int
	Length int
}

func SumPartNumbers(input []string) int {
	schema := newSchema(input)
	parts := parsePartsWithMath(&schema)
	return calculateNumberTotal(parts, &schema)
}

func parsePartsWithRegex(input []string) []EnginePart {
	var parts []EnginePart
	for row, line := range input {
		finder := regexp.MustCompile(`[0-9]+`)
		partStrings := finder.FindAllString(line, -1)
		partStringIndexes := finder.FindAllStringIndex(line, -1)

		for i, partString := range partStrings {
			number, _ := strconv.Atoi(partString)
			column := partStringIndexes[i][0]
			length := len([]rune(partString))
			part := EnginePart{number, row, column, length}
			parts = append(parts, part)
		}
	}

	return parts
}

func parsePartsWithMath(s *Schema) []EnginePart {
	var parts []EnginePart
	for row, runes := range s.Content {
		for column := 0; column < s.Columns; {
			digit, ok := kit.RuneToInt(runes[column])
			if !ok {
				column++
				continue
			}

			digits := []int{digit}
			for next := column + 1; next < s.Columns; next++ {
				digit, ok := kit.RuneToInt(runes[next])
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

			part := EnginePart{number, row, column, length}
			parts = append(parts, part)

			column += length
		}
	}

	return parts
}

func calculateNumberTotal(parts []EnginePart, schema *Schema) int {
	total := 0
	for _, part := range parts {
		neighbors := findNeighbors(part, schema)
		for _, neighbor := range neighbors {
			if neighbor.Symbol == '.' {
				continue
			}

			if '0' <= neighbor.Symbol && neighbor.Symbol <= '9' {
				continue
			}

			total += part.Number
			break
		}
	}

	return total
}
