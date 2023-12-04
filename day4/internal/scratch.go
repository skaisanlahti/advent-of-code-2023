package internal

import (
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type scratchCard struct {
	wins  int
	count int
}

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

func parseCards(input []string) []scratchCard {
	var cards []scratchCard
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

		card := scratchCard{wins, 1}
		cards = append(cards, card)
	}

	return cards
}

func countPoints(cards []scratchCard) int {
	total := 0
	for _, card := range cards {
		points := 0
		for i := 0; i < card.wins; i++ {
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

func countCards(cards []scratchCard) int {
	total := 0
	for index, card := range cards {
		total += card.count

		start := index + 1
		end := start + card.wins
		for copy := 0; copy < card.count; copy++ {
			for target := start; target < end; target++ {
				cards[target].count += 1
			}
		}
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
