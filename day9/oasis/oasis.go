package oasis

import (
	"strconv"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func PredictNextValue(input []string, reverse bool) int {
	histories := parseHistories(input)
	if reverse {
		for _, history := range histories {
			kit.ReverseInts(history)
		}
	}

	sum := 0
	for _, history := range histories {
		sum += recursiveExtrapolate(history)
	}

	return sum
}

func parseHistories(lines []string) [][]int {
	histories := [][]int{}
	for _, line := range lines {
		history := []int{}
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				continue
			}

			history = append(history, value)
		}

		histories = append(histories, history)
	}

	return histories
}

func isFinal(history []int) bool {
	isFinal := true
	for i := 0; i < len(history); i++ {
		if history[i] != 0 {
			isFinal = false
			break
		}
	}

	return isFinal
}

func recursiveExtrapolate(history []int) int {
	if isFinal(history) {
		return 0
	}

	deltas := []int{}
	for i := 1; i < len(history); i++ {
		deltas = append(deltas, history[i]-history[i-1])
	}

	return history[len(history)-1] + recursiveExtrapolate(deltas)
}

func loopExtrapolate(history []int) int {
	current := history
	values := [][]int{current}
	for !isFinal(current) {
		deltas := []int{}
		for i := 1; i < len(current); i++ {
			deltas = append(deltas, current[i]-current[i-1])
		}

		values = append(values, deltas)
		current = deltas
	}

	current = append(current, 0)
	values[len(values)-1] = current

	for i := len(values) - 2; i >= 0; i-- {
		this := values[i]
		last := values[i+1]
		this = append(this, this[len(this)-1]+last[len(last)-1])
		values[i] = this
	}

	return values[0][len(values[0])-1]
}
