package internal

import (
	"fmt"
)

func SumGearRatios(input []string) int {
	schema := newSchema(input)
	parts := findEngineParts(input)
	gears := mapPartsToGears(parts, &schema)
	return calculateGearRatios(gears)
}

func mapPartsToGears(parts []enginePart, schema *schema) map[string][]int {
	gears := map[string][]int{}
	for _, part := range parts {
		neighbors := findNeighbors(part, schema)
		for _, neighbor := range neighbors {
			ok := isGear(neighbor.char)
			if ok {
				key := fmt.Sprintf("%d.%d", neighbor.row, neighbor.column)
				numbers, found := gears[key]
				if !found {
					numbers = []int{}
				}

				gears[key] = append(numbers, part.number)
			}
		}
	}

	return gears
}

func calculateGearRatios(gears map[string][]int) int {
	total := 0
	for _, numbers := range gears {
		if len(numbers) == 2 {
			ratio := numbers[0] * numbers[1]
			total += ratio
		}
	}

	return total
}

func isGear(r rune) bool {
	if r == '*' {
		return true
	}

	return false
}
