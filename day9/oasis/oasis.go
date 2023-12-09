package oasis

import (
	"strconv"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func PredictNextValue(input []string, reverse bool) int {
	histories := parseHistories(input)
	if reverse {
		histories = reverseHistories(histories)
	}

	sum := 0
	for _, history := range histories {
		prediction := extrapolate(history)
		sum += prediction
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
		}
	}

	return isFinal
}

func extrapolate(history []int) int {
	if isFinal(history) {
		return 0
	}

	differences := []int{}
	for i := 1; i < len(history); i++ {
		differences = append(differences, history[i]-history[i-1])
	}

	next := extrapolate(differences)
	return history[len(history)-1] + next
}

func reverseHistories(histories [][]int) [][]int {
	reversed := [][]int{}
	for _, history := range histories {
		reversed = append(reversed, kit.ReverseSlice(history))
	}

	return reversed
}
