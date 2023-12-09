package kit

import "math"

func NumbersFromString(line string) []int {
	runes := []rune(line)
	return NumbersFromRunes(runes)
}

func NumbersFromRunes(runes []rune) []int {
	numbers := []int{}
	for i := 0; i < len(runes); {
		digit, ok := RuneToInt(runes[i])
		if !ok {
			i++
			continue
		}

		digits := []int{digit}
		for next := i + 1; next < len(runes); next++ {
			digit, ok := RuneToInt(runes[next])
			if !ok {
				break
			}

			digits = append(digits, digit)
		}

		length := len(digits)
		number := 0
		for i, m := 0, length-1; i < length; i, m = i+1, m-1 {
			number += digits[i] * int(math.Pow10(m))
		}

		numbers = append(numbers, number)

		i += length
	}

	return numbers
}

func RuneToInt(r rune) (int, bool) {
	if r < '0' || '9' < r {
		return 0, false
	}

	return int(r - '0'), true
}
