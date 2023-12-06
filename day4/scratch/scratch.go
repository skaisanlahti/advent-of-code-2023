package scratch

import (
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func CountPoints(input []string) int {
	wins := parseWinCounts(input)
	return countPoints(wins)
}

func CountCards(input []string) int {
	wins := parseWinCounts(input)
	return countCards(wins)
}

func parseWinCounts(input []string) []int {
	winCounts := []int{}
	for _, line := range input {
		numbers := strings.Split(line, ":")[1]
		splits := strings.Split(numbers, "|")
		winningNumbers := kit.NumbersFromString(splits[0])
		scratchedNumbers := kit.NumbersFromString(splits[1])

		wins := 0
		for _, scratchedNumber := range scratchedNumbers {
			for _, winningNumber := range winningNumbers {
				if scratchedNumber == winningNumber {
					wins++
					break
				}
			}
		}

		winCounts = append(winCounts, wins)
	}

	return winCounts
}

func countPoints(winCounts []int) int {
	totalPoints := 0
	for _, wins := range winCounts {
		points := 0
		for i := 0; i < wins; i++ {
			switch points {
			case 0:
				points++
			default:
				points *= 2
			}
		}

		totalPoints += points
	}

	return totalPoints
}

func countCards(winCounts []int) int {
	cardCounts := map[int]int{}
	totalCards := 0
	for card, wins := range winCounts {
		cardCounts[card] += 1
		cardCount := cardCounts[card]
		totalCards += cardCount

		start := card + 1
		end := start + wins
		for i := 0; i < cardCount; i++ {
			for target := start; target < end; target++ {
				cardCounts[target] += 1
			}
		}
	}

	return totalCards
}
