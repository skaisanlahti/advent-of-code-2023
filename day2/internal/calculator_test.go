package internal

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type powerCase struct {
	input    []string
	expected int
}

var powerCases = []powerCase{
	{kit.ReadInputFile("../input/sample.txt"), 2286},
	{kit.ReadInputFile("../input/data.txt"), 83707},
}

func TestSumGamePowers(t *testing.T) {
	for i, testCase := range powerCases {
		result := SumGamePowers(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkSumGamePowers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumGamePowers(powerCases[1].input)
	}
}
