package kit

import "math"

func TernarySearch(function func(float64) float64, left, right, precision float64) float64 {
	for math.Abs(right-left) >= precision {
		nextLeft := left + (right-left)/3
		nextRight := right - (right-left)/3

		leftResult := function(nextLeft)
		rightResult := function(nextRight)

		if leftResult < rightResult {
			left = nextLeft
		} else {
			right = nextRight
		}
	}
	return (left + right) / 2
}
