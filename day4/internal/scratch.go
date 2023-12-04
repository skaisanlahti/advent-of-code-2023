package internal

import (
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func CountPoints(input []string) int {
	cards := parseCards(input)
	total := countPoints(cards)
	return total
}

func CountCopies(input []string) int {
	cards := parseCards(input)
	total := countCopies(cards)
	return total
}

func parseCards(input []string) []int {
	cards := []int{} // keeps count of wins
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

func countCopies(cards []int) int {
	copies := map[int]int{} // maps index to count
	total := 0
	for index, wins := range cards {
		copies[index] += 1
		count := copies[index]
		total += count

		start := index + 1
		end := start + wins
		for copy := 0; copy < count; copy++ {
			for target := start; target < end; target++ {
				copies[target] += 1
			}
		}
	}

	return total
}
