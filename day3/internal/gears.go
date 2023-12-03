package internal

import (
	"fmt"
)

func SumGearRatios(input []string) int {
	schema := newSchema(input)
	parts := parsePartsWithMath(&schema)
	gears := mapPartsToGears(parts, &schema)
	return calculateGearRatios(gears)
}

func mapPartsToGears(parts []enginePart, schema *schema) map[string][]int {
	gears := map[string][]int{}
	for _, part := range parts {
		neighbors := findNeighbors(part, schema)
		for _, neighbor := range neighbors {
			if neighbor.symbol != '*' {
				continue
			}

			key := fmt.Sprintf("%d.%d", neighbor.row, neighbor.column)
			numbers, ok := gears[key]
			if !ok {
				numbers = []int{}
			}

			gears[key] = append(numbers, part.number)
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
