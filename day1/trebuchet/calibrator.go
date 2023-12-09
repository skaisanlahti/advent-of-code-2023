package trebuchet

import (
	"strconv"
	"strings"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

func SumCalibrationValues(input []string, checkStringPatterns bool) int {
	total := 0
	for _, line := range input {
		first, ok := findDigit(line, checkStringPatterns, false)
		if !ok {
			continue // line contains no valid digits, skip checking in reverse order
		}

		last, _ := findDigit(line, checkStringPatterns, true)
		value, _ := strconv.Atoi(first + last)
		total += value
	}

	return total
}

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func findDigit(line string, checkStringPatterns bool, reverseOrder bool) (string, bool) {
	// find digit either from start or end of line
	if reverseOrder {
		line = kit.ReverseString(line)
	}

	// find digit based on number value
	matches := map[string]int{}
	for index, char := range line {
		digit := string(char)
		value, err := strconv.Atoi(digit)
		if err != nil || value == 0 {
			continue
		}

		// if string patterns are not included in search, we can exit on first valid digit
		if !checkStringPatterns {
			return digit, true
		}

		// save index to compare with string pattern indexes
		matches[digit] = index
		break
	}

	// find digits with string patterns
	for key := range digits {
		pattern := key
		if reverseOrder {
			pattern = kit.ReverseString(key)
		}

		index := strings.Index(line, pattern)
		if index != -1 {
			matches[key] = index
		}
	}

	// find which pattern had the lowest index
	lowest := len(line)
	pattern := ""
	for key, index := range matches {
		if index < lowest {
			lowest = index
			pattern = key
		}
	}

	// return pattern if it's a digit or find the digit using the pattern
	_, err := strconv.Atoi(pattern)
	if err != nil {
		digit, ok := digits[pattern]
		return digit, ok
	}

	return pattern, true
}
