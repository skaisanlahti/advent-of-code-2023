package scratch

import (
	"testing"

	"github.com/skaisanlahti/advent-of-code-2023/kit"
)

type testCase struct {
	input    []string
	expected int
}

var testCases = []testCase{
	{kit.ReadInputFile("../input/sample.txt"), 13},
	{kit.ReadInputFile("../input/data.txt"), 15268},
	{kit.ReadInputFile("../input/sample.txt"), 30},
	{kit.ReadInputFile("../input/data.txt"), 6283755},
}

func TestCountPoints(t *testing.T) {
	for i, testCase := range testCases[0:1] {
		result := CountPoints(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkCountPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPoints(testCases[1].input)
	}
}

func TestCountCards(t *testing.T) {
	for i, testCase := range testCases[2:3] {
		result := CountCards(testCase.input)
		if result != testCase.expected {
			t.Errorf("Test case %d expected %d but result was %d.", i+1, testCase.expected, result)
		}
	}
}

func BenchmarkCountCards(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountCards(testCases[3].input)
	}
}
