package calibrator

import (
	"log"
	"strconv"
	"strings"
)

func Calibrate(input []string, checkStrings bool) int {
	log.Println("Calibration:")
	log.Println("---")
	var values []int
	for _, line := range input {
		firstDigit := findDigit(line, checkStrings, false)
		lastDigit := findDigit(line, checkStrings, true)

		value, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Panicln(err.Error())
		}

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

func findDigit(line string, checkStrings bool, reverseOrder bool) string {
	if reverseOrder {
		line = reverseString(line)
	}

	matches := map[string]int{}
	for index, char := range line {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		matches[string(char)] = index
		break
	}

	for key := range patterns {
		if checkStrings {
			localKey := key
			if reverseOrder {
				localKey = reverseString(key)
			}

			keyIndex := strings.Index(line, localKey)
			if keyIndex != -1 {
				matches[key] = keyIndex
			}
		}
	}

	firstIndex := len(line)
	var candidate string
	for key, value := range matches {
		if value < firstIndex {
			firstIndex = value
			candidate = key
		}
	}

	_, err := strconv.Atoi(candidate)
	if err != nil {
		conversion, ok := patterns[candidate]
		if !ok {
			log.Panicln("Pattern conversion not found.")
		}

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
