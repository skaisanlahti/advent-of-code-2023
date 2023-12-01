package calibrator

import (
	"log"
	"strconv"
	"strings"
)

var patterns = map[string]string{
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

func Calibrate(input []string, checkStrings bool) int {
	log.Println("Calibration:")
	log.Println("---")
	var values []int
	for _, line := range input {
		value := findValue(line, checkStrings)
		values = append(values, value)
		log.Println(value)
	}

	sum := 0
	for _, number := range values {
		sum += number
	}
	log.Println("---")
	log.Printf("=%d", sum)
	log.Println("---")
	return sum
}

func findValue(line string, checkStrings bool) int {
	first := findFirstDigit(line, checkStrings)
	last := findLastDigit(line, checkStrings)
	value, err := strconv.Atoi(first + last)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return value
}

func findFirstDigit(line string, checkStrings bool) string {
	results := map[string]int{}
	for key, value := range patterns {
		valueIndex := strings.Index(line, value)
		if valueIndex != -1 {
			results[value] = valueIndex
		}

		if checkStrings {
			keyIndex := strings.Index(line, key)
			if keyIndex != -1 {
				results[key] = keyIndex
			}
		}
	}

	firstIndex := len(line)
	var candidate string
	for key, value := range results {
		if value < firstIndex {
			firstIndex = value
			candidate = key
		}
	}

	_, err := strconv.Atoi(candidate)
	if err != nil {
		conversion, _ := patterns[candidate]
		_, err = strconv.Atoi(conversion)
		if err != nil {
			log.Panicln(err)
		}
		return conversion
	}

	return candidate
}

func findLastDigit(line string, checkStrings bool) string {
	revLine := reverseString(line)
	results := map[string]int{}
	for key, value := range patterns {
		valueIndex := strings.Index(revLine, value)
		if valueIndex != -1 {
			results[value] = valueIndex
		}

		if checkStrings {
			revKey := reverseString(key)
			keyIndex := strings.Index(revLine, revKey)
			if keyIndex != -1 {
				results[key] = keyIndex + len(key) - 1
			}
		}

	}

	firstIndex := len(line)
	var candidate string
	for key, value := range results {
		if value < firstIndex {
			firstIndex = value
			candidate = key
		}
	}

	_, err := strconv.Atoi(candidate)
	if err != nil {
		conversion, _ := patterns[candidate]
		_, err = strconv.Atoi(conversion)
		if err != nil {
			log.Panicln(err)
		}
		return conversion
	}

	return candidate
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
