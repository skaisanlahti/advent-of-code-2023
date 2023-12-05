package internal

import (
	"math"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func FindLowestLocationValue(input []string) int {
	seeds := []int{}
	stages := [][][]int{}
	stageIndex := -1
	for i, line := range input {
		if i == 0 {
			seeds = kit.NumbersFromString(line)
			continue
		}

		if strings.Contains(line, ":") {
			stageIndex++
			stages = append(stages, [][]int{})
			continue
		}

		numbers := kit.NumbersFromString(line)
		if len(numbers) > 0 {
			transformation := stages[stageIndex]
			transformation = append(transformation, numbers)
			stages[stageIndex] = transformation
		}
	}

	values := seeds
	for _, stage := range stages {
		nextValues := []int{}
		for _, value := range values {
			processed := false
			for _, transformation := range stage {
				destination := transformation[0]
				source := transformation[1]
				length := transformation[2]

				// if value is in the transformation range apply changes and break
				if source <= value && value < source+length {
					nextValue := value + destination - source
					nextValues = append(nextValues, nextValue)
					processed = true
					break
				}
			}

			// value wasn't transformed, add to next stage without changes
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

func FindLowestLocationValueInRange(input []string) int {
	seeds := [][]int{}    // start end
	stages := [][][]int{} // destination source length
	stageIndex := -1
	for i, line := range input {
		if i == 0 {
			numbers := kit.NumbersFromString(line)
			for i := 0; i < len(numbers); i += 2 {
				seeds = append(seeds, []int{numbers[i], numbers[i] + numbers[i+1]})
			}
			continue
		}

		if strings.Contains(line, ":") {
			stageIndex++
			stages = append(stages, [][]int{})
			continue
		}

		numbers := kit.NumbersFromString(line)
		if len(numbers) > 0 {
			transformation := stages[stageIndex]
			transformation = append(transformation, numbers)
			stages[stageIndex] = transformation
		}
	}

	values := seeds // working set
	// handle one transformation stage at a time
	for _, stage := range stages {
		// transform values from values to nextValues set
		nextValues := [][]int{}
		for len(values) > 0 {
			// take out value from values
			value := values[0]
			values = values[1:]

			start := value[0]
			end := value[1]
			processed := false
			for _, transformation := range stage {
				destination := transformation[0]
				source := transformation[1]
				length := transformation[2]

				// check which parts of the range overlap with transformation range
				overlapStart := math.Max(float64(start), float64(source))
				overlapEnd := math.Min(float64(end), float64(source+length))
				if overlapStart < overlapEnd {
					// apply offset to the overlapping part
					offset := destination - source
					nextStart := int(overlapStart) + offset
					nextEnd := int(overlapEnd) + offset
					nextValues = append(nextValues, []int{nextStart, nextEnd})

					// add values that didn't overlap with this transformation range back to values
					// to check if they overlap with over transformation ranges
					if int(overlapStart) > start {
						values = append(values, []int{start, int(overlapStart)})
					}
					if end > int(overlapEnd) {
						values = append(values, []int{int(overlapEnd), end})
					}

					// values only need to be processed once so break to skip rest of ranges
					processed = true
					break
				}
			}

			// range didn't overlap with anything, add to next stage
			if !processed {
				nextValues = append(nextValues, []int{start, end})
			}
		}

		// all values have been transformed, assign to values again for next stage
		// as we loop to the next block of transformations
		values = nextValues
	}

	found := math.MaxInt
	for _, value := range values {
		found = int(math.Min(float64(found), float64(value[0])))
	}

	return found
}
