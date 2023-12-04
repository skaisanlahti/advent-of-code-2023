package internal

import (
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func SumCardPoints(input []string) int {
	cards := parseCards(input)
	total := countPoints(cards)
	return total
}

func SumCardCount(input []string) int {
	cards := parseCards(input)
	total := countCards(cards)
	return total
}

func parseCards(input []string) []int {
	cards := []int{}
	for _, line := range input {
		numbers := strings.Split(line, ":")[1]
		splits := strings.Split(numbers, "|")
		winningNumbers := kit.NumbersFromString(splits[0])
		scratchedNumbers := kit.NumbersFromString(splits[1])

		wins := 0
		for _, scratched := range scratchedNumbers {
			for _, winning := range winningNumbers {
				if scratched == winning {
					wins++
					break
				}
			}
		}

		cards = append(cards, wins)
	}

	return cards
}

func countPoints(cards []int) int {
	total := 0
	for _, wins := range cards {
		points := 0
		for i := 0; i < wins; i++ {
			switch points {
			case 0:
				points++
			default:
				points *= 2
			}
		}

		total += points
	}

	return total
}

func countCards(cards []int) int {
	counts := map[int]int{}
	total := 0
	for index, wins := range cards {
		counts[index] += 1
		count := counts[index]
		total += count

		start := index + 1
		end := start + wins
		for copy := 0; copy < count; copy++ {
			for target := start; target < end; target++ {
				counts[target] += 1
			}
		}
	}

	return total
}

func CountPointsSinglePass(input []string) int {
	total := 0
	for _, line := range input {
		// parse numbers
		numbers := strings.Split(line, ":")[1]
		splits := strings.Split(numbers, "|")
		winningNumbers := kit.NumbersFromString(splits[0])
		scratchedNumbers := kit.NumbersFromString(splits[1])

		// check wins
		win := map[int]int{}
		for _, number := range winningNumbers {
			win[number] = number
		}

		points := 0
		for _, number := range scratchedNumbers {
			_, ok := win[number]
			if ok {
				switch points {
				case 0:
					points++
				default:
					points *= 2
				}
			}
		}

		total += points
	}

	return total
}

func CountCardsSinglePass(input []string) int {
	// map card index to count
	counts := map[int]int{}
	total := 0
	for index, line := range input {
		// parse numbers
		numbers := strings.Split(line, ":")[1]
		splits := strings.Split(numbers, "|")
		winningNumbers := kit.NumbersFromString(splits[0])
		scratchedNumbers := kit.NumbersFromString(splits[1])

		// check wins
		wins := 0
		for _, scratched := range scratchedNumbers {
			for _, winning := range winningNumbers {
				if scratched == winning {
					wins++
					break
				}
			}
		}

		// count copies
		counts[index] += 1
		count := counts[index]
		total += count

		// add copies of further cards based on wins
		start := index + 1
		end := start + wins
		for copy := 0; copy < count; copy++ {
			for target := start; target < end; target++ {
				counts[target] += 1
			}
		}
	}

	return total
}
