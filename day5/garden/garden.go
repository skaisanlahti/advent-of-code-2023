package garden

import (
	"math"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func FindLowestLocationValue(input []string) int {
	values := kit.NumbersFromString(input[0])
	stages := parseMappingStages(input[1:])
	for _, mappers := range stages {
		nextValues := []int{}
		for _, value := range values {
			processed := false
			for _, mapper := range mappers {
				if mapper.Source <= value && value < mapper.Source+mapper.Length {
					nextValue := value + mapper.Destination - mapper.Source
					nextValues = append(nextValues, nextValue)
					processed = true
					break
				}
			}

			if !processed {
				nextValues = append(nextValues, value)
			}
		}

		values = nextValues
	}

	found := math.MaxInt
	for _, value := range values {
		found = int(math.Min(float64(found), float64(value)))
	}

	return found
}

type ValueRange struct {
	Start int
	End   int
}

func FindLowestLocationRange(input []string) int {
	valueRanges := []ValueRange{}
	seedRanges := kit.NumbersFromString(input[0])
	for i := 0; i < len(seedRanges); i += 2 {
		start := seedRanges[i]
		end := seedRanges[i] + seedRanges[i+1]
		valueRanges = append(valueRanges, ValueRange{start, end})
	}

	stages := parseMappingStages(input[1:])
	for _, mappers := range stages {
		nextValueRanges := []ValueRange{}
		for len(valueRanges) > 0 {
			valueRange := valueRanges[0]
			valueRanges = valueRanges[1:]
			processed := false
			for _, mapper := range mappers {
				overlapStart := int(math.Max(float64(valueRange.Start), float64(mapper.Source)))
				overlapEnd := int(math.Min(float64(valueRange.End), float64(mapper.Source+mapper.Length)))
				if overlapStart < overlapEnd {
					nextStart := overlapStart + mapper.Destination - mapper.Source
					nextEnd := overlapEnd + mapper.Destination - mapper.Source
					nextValueRanges = append(nextValueRanges, ValueRange{nextStart, nextEnd})

					if overlapStart > valueRange.Start {
						valueRanges = append(valueRanges, ValueRange{valueRange.Start, overlapStart})
					}

					if valueRange.End > overlapEnd {
						valueRanges = append(valueRanges, ValueRange{overlapEnd, valueRange.End})
					}

					processed = true
					break
				}
			}

			if !processed {
				nextValueRanges = append(nextValueRanges, valueRange)
			}
		}

		valueRanges = nextValueRanges
	}

	found := math.MaxInt
	for _, valueRange := range valueRanges {
		found = int(math.Min(float64(found), float64(valueRange.Start)))
	}

	return found
}

type Mapper struct {
	Destination int
	Source      int
	Length      int
}

func parseMappingStages(input []string) [][]Mapper {
	stages := [][]Mapper{}
	stageIndex := -1
	for _, line := range input {
		if strings.Contains(line, ":") {
			stageIndex++
			stages = append(stages, []Mapper{})
			continue
		}

		mapper := kit.NumbersFromString(line)
		if len(mapper) > 0 {
			mappers := stages[stageIndex]
			mappers = append(mappers, Mapper{mapper[0], mapper[1], mapper[2]})
			stages[stageIndex] = mappers
		}
	}

	return stages
}
