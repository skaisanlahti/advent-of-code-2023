package internal

import (
	"math"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type Mapper struct {
	destination int
	source      int
	length      int
}

func FindLowestLocationValue(input []string) int {
	values := []int{}
	stages := [][]Mapper{}
	stageIndex := -1
	for i, line := range input {
		if i == 0 {
			seeds := kit.NumbersFromString(line)
			values = seeds
			continue
		}

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

	for _, mappers := range stages {
		nextValues := []int{}
		for _, value := range values {
			processed := false
			for _, mapper := range mappers {
				if mapper.source <= value && value < mapper.source+mapper.length {
					nextValue := value + mapper.destination - mapper.source
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
	start int
	end   int
}

func FindLowestLocationRange(input []string) int {
	valueRanges := []ValueRange{}
	stages := [][]Mapper{}
	stageIndex := -1
	for i, line := range input {
		if i == 0 {
			seedRanges := kit.NumbersFromString(line)
			for i := 0; i < len(seedRanges); i += 2 {
				start := seedRanges[i]
				end := seedRanges[i] + seedRanges[i+1]
				valueRanges = append(valueRanges, ValueRange{start, end})
			}
			continue
		}

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

	for _, mappers := range stages {
		nextValueRanges := []ValueRange{}
		for len(valueRanges) > 0 {
			valueRange := valueRanges[0]
			valueRanges = valueRanges[1:]
			processed := false
			for _, mapper := range mappers {
				overlapStart := int(math.Max(float64(valueRange.start), float64(mapper.source)))
				overlapEnd := int(math.Min(float64(valueRange.end), float64(mapper.source+mapper.length)))
				if overlapStart < overlapEnd {
					nextStart := overlapStart + mapper.destination - mapper.source
					nextEnd := overlapEnd + mapper.destination - mapper.source
					nextValueRanges = append(nextValueRanges, ValueRange{nextStart, nextEnd})

					if overlapStart > valueRange.start {
						valueRanges = append(valueRanges, ValueRange{valueRange.start, overlapStart})
					}

					if valueRange.end > overlapEnd {
						valueRanges = append(valueRanges, ValueRange{overlapEnd, valueRange.end})
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
		found = int(math.Min(float64(found), float64(valueRange.start)))
	}

	return found
}
